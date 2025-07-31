package fixture

import (
	"context"
	"fmt"
	"medblogers_base/e2e/matcher"
	"medblogers_base/internal/pkg/postgres"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var SetupDatabase = func(ctx context.Context, seed int64) {

	seed += GinkgoRandomSeed()

	By("Создание базы данных")
	conn, dbName, err := conn(ctx, seed, GinkgoParallelProcess())
	Expect(err).NotTo(HaveOccurred())
	_, err = conn.Exec(ctx, fmt.Sprintf("create database %s", dbName))
	Expect(err).NotTo(HaveOccurred())
	Expect(conn.Close(ctx)).NotTo(HaveOccurred())

	By("Применение миграций")
	cmd := exec.Command("make", "migrations-e2e-up")
	cmd.Env = append(os.Environ(), fmt.Sprintf("E2E_DB_NAME=%s", dbName))
	_, f, _, _ := runtime.Caller(0)
	// Получаем рут директорию
	cmd.Dir = strings.ReplaceAll(f, "/e2e/fixture/prepare_postgres.go", "")

	Expect(err).NotTo(HaveOccurred())
	_, err = cmd.Output()
	Expect(err).NotTo(matcher.HaveOccurred())

	DeferCleanup(func(ctx context.Context) {
		TearDownDatabase(ctx, seed)
	})
}

var SetupPoolConnections = func(ctx context.Context, seed int64) postgres.PoolWrapper {
	By("Подготовка подключения к БД")

	seed += GinkgoRandomSeed()

	dbName := fmt.Sprintf("temp_db_%d_%d", seed, GinkgoParallelProcess())

	DSN := fmt.Sprintf("user=postgres dbname=%s host=localhost port=5432 sslmode=disable", dbName)

	poolConfig, err := pgxpool.ParseConfig(DSN)
	Expect(err).NotTo(HaveOccurred())

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)

	Expect(err).NotTo(HaveOccurred())

	DeferCleanup(func(ctx context.Context) {
		if pool != nil {
			pool.Close()
		}
	})

	return postgres.NewPoolWrapper(pool)
}

var TearDownDatabase = func(ctx context.Context, seed int64) {
	By("Удаление базы")
	conn, dbName, err := conn(ctx, seed, GinkgoParallelProcess())
	Expect(err).NotTo(HaveOccurred())
	_, err = conn.Exec(ctx, fmt.Sprintf("REVOKE CONNECT ON DATABASE %s FROM public;", dbName))
	Expect(err).NotTo(HaveOccurred())
	_, err = conn.Exec(ctx, fmt.Sprintf("select pg_terminate_backend(pg_stat_activity.pid) from pg_stat_activity where pg_stat_activity.datname = '%s';", dbName))
	Expect(err).NotTo(HaveOccurred())
	_, err = conn.Exec(ctx, fmt.Sprintf("drop database if exists %s", dbName))
	Expect(err).NotTo(HaveOccurred())
	Expect(conn.Close(ctx)).NotTo(HaveOccurred())
}

func conn(ctx context.Context, seed int64, process int) (*pgx.Conn, string, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres@localhost")
	dbName := fmt.Sprintf("temp_db_%d_%d", seed, process)

	return conn, dbName, err
}
