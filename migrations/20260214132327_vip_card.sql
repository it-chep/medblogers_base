-- +goose Up
-- +goose StatementBegin

-- ВИП - карточка
create table if not exists vip_card
(
    doctor_id              bigint primary key,

    can_barter             bool,   -- Есть бартер
    can_buy_advertising    bool,   -- Куплю рекламу
    can_sell_advertising   bool,   -- Продам рекламу

    short_message          text,   -- Свободный статус/послание кому-то
    advertising_price_from bigint, -- цена рекламы от
    blog_info              text,   -- Расширенное описание: сколько лет блогу, как часто ведёт тг-канал

    -- Настройка вип карточки
    end_date               timestamp
);

alter table docstar_site_doctor
    add column is_vip bool default false;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table vip_card;

alter table docstar_site_doctor
    drop column is_vip;
-- +goose StatementEnd
