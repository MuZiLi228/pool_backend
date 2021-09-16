
-- 修改字段为浮点类型   小数保留4位
alter table shareholder_daily_income alter column cashable_income type decimal(100,5) using cashable_income::decimal(100,5);
alter table shareholder_daily_income alter column the_nominal_income type decimal(100,5) using the_nominal_income::decimal(100,5);

alter table withdrawal alter column amount type decimal(100,5) using amount::decimal(100,5);

alter table fil_pool alter column miner_balance type decimal(100,5) using miner_balance::decimal(100,5);
alter table fil_pool alter column miner_available_balance type decimal(100,5) using miner_available_balance::decimal(100,5);
alter table fil_pool alter column sector_size type decimal(100,5) using sector_size::decimal(100,5);


alter table fil_pool_daily_income alter column assigned_val type decimal(100,5) using assigned_val::decimal(100,5);
alter table fil_pool_daily_income alter column freed type decimal(100,5) using freed::decimal(100,5);
alter table fil_pool_daily_income alter column last_time_val type decimal(100,5) using last_time_val::decimal(100,5);
alter table fil_pool_daily_income alter column today_val type decimal(100,5) using today_val::decimal(100,5);




