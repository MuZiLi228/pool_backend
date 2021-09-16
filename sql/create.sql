-- 
-- 创建用户test_user并设置password为test_password
create user fil_pool with password 'fUz3iBspp7tX8oNRy2fs';

-- 创建database test_db设置owner为test_user
create database fil_pool with owner fil_pool;

-- 切换到新创建的test_db数据库
\c fil_pool fil_pool;


-- 股东表
CREATE TABLE IF NOT EXISTS shareholder (
    id text not null default '',
    mobile text not null default '',
    role text not null default '',
    login_pwd text not null default '',
    withdrawal_pwd text not null default '',
    income int not null default 0,
    is_enable boolean not null default true,
    withdrawal_limit int not null default 0,
    percent_three_shareholder_id text not null default '',
    percent_five_shareholder_id text not null default '',
    recommend_shareholder_id text not null default '',
    recommend_code text not null default '',
    recommend_num int not null default 0,
    recommend_allocation_ratio int not null default 0,
    recent_withdrawal_account text not null default '',
    fil_pool_num int not null default 0,
    create_at timestamp with time zone not null default current_timestamp,
    CONSTRAINT pk_shareholder_id PRIMARY KEY (id)
);
comment on column shareholder.income is '总收益';
comment on column shareholder.login_pwd is '登录密码';
comment on column shareholder.withdrawal_pwd is '提现密码';
comment on column shareholder.recommend_shareholder_id is '推荐者id';
comment on column shareholder.recommend_code is '推荐码';
comment on column shareholder.withdrawal_limit is '可提现余额';
comment on column shareholder.recommend_allocation_ratio is '推广分配比例';
comment on column shareholder.recommend_num is '推广人数';
comment on column shareholder.fil_pool_num is '矿池数量';


-- 矿池表
CREATE TABLE IF NOT EXISTS fil_pool (
    id text not null default '',
    name text not null default '',
    miner text not null default '',
    miner_balance int not null default 0,
    role text not null default '',
    miner_available_balance int not null default 0,
    sector_size int not null default 0,
    effective_computing_power text not null default '',
    original_computing_power text not null default '',
    shareholders_num int not null default 0,
    shareholder_distribution_ratio text[] not null default array[]::text[],
    node_id text not null default '',
    create_at timestamp with time zone not null default current_timestamp,
    CONSTRAINT pk_fil_pool_id PRIMARY KEY (id)
);
comment on column fil_pool.miner is '矿工账号';
comment on column fil_pool.miner_balance is '账户余额';
comment on column fil_pool.miner_available_balance is '可用余额';
comment on column fil_pool.sector_size is '扇区大小';
comment on column fil_pool.effective_computing_power is '有效算力';
comment on column fil_pool.original_computing_power is '原值算力';
comment on column fil_pool.shareholders_num is '股东人数';
comment on column fil_pool.shareholder_distribution_ratio is '股东分配比例 json格式{A:20, B:50, C:30}';
comment on column fil_pool.node_id is '节点id';


-- 股东每日收益
CREATE TABLE IF NOT EXISTS shareholder_daily_income (
    id text not null default '',
    shareholder_id text not null default '',
    income_type text not null default '',
    fil_pool_daily_income_id text not null default '',
    cashable_income int not null default 0,
    the_nominal_income int not null default 0,
    create_at timestamp with time zone not null default current_timestamp,
    CONSTRAINT pk_shareholder_daily_income_id PRIMARY KEY (id)
);




-- 矿池每日收益 
CREATE TABLE IF NOT EXISTS fil_pool_daily_income (
    id text not null default '',
    fil_pool_id text not null default '',
    is_allocated boolean not null default false,
    assigned_val decimal(100,4) not null default 0.0,
    freed decimal(100,4) not null default 0.0,
    last_time_val decimal(100,4) not null default 0.0,
    today_val decimal(100,4) not null default 0.0,
    create_at timestamp with time zone not null default current_timestamp,
    CONSTRAINT pk_fil_pool_daily_income_id PRIMARY KEY (id)
);



--  股东矿池比例
CREATE TABLE IF NOT EXISTS fil_pool_ratio(
    fil_pool_id text not null default '',
    shareholder_id text not null default '',
    proportion_of_shares int not null default 0,
    create_at timestamp with time zone not null default current_timestamp,
    end_at timestamp with time zone not null default current_timestamp,
    update_at timestamp with time zone not null default current_timestamp
);

-- 更新后的矿池比例 用于结算分配
CREATE TABLE IF NOT EXISTS updated_fil_pool_ratio(
    fil_pool_id text not null default '',
    shareholder_id text not null default '',
    proportion_of_shares int not null default 0,
    create_at timestamp with time zone not null default current_timestamp,
    end_at timestamp with time zone not null default current_timestamp,
    update_at timestamp with time zone not null default current_timestamp
);


--提现申请
CREATE TABLE IF NOT EXISTS withdrawal(
    id text not null default '',
    shareholder_id text not null default '',
    amount int not null default 0,
    state int not null default 0, 
    content text not null default '',
    with_log text not null default '',
    hash text not null default '',
    create_at timestamp with time zone not null default current_timestamp,
    end_at timestamp with time zone not null default current_timestamp,
    CONSTRAINT pk_withdrawal_id PRIMARY KEY (id)
);

