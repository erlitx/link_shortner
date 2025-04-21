BEGIN;

CREATE TABLE wb_orders (
    id               SERIAL PRIMARY KEY,         -- Auto-incremented ID
    date            TIMESTAMP WITH TIME ZONE,    -- Order date
    last_change_date TIMESTAMP WITH TIME ZONE,   -- Last modified date
    warehouse_name   TEXT,                       -- Warehouse name
    warehouse_type   TEXT,                       -- Warehouse type
    country_name     TEXT,                       -- Country
    oblast_okrug_name TEXT,                      -- Federal district
    region_name      TEXT,                       -- Region
    supplier_article TEXT,                       -- Supplier's article
    nm_id            INTEGER,                    -- Product ID
    barcode          TEXT,                       -- Barcode
    category         TEXT,                       -- Product category
    subject          TEXT,                       -- Subject
    brand            TEXT,                       -- Brand
    tech_size        TEXT,                       -- Technical size
    income_id        INTEGER,                    -- Income ID
    is_supply        BOOLEAN,                    -- Is it a supply order?
    is_realization   BOOLEAN,                    -- Is it a realization order?
    total_price      NUMERIC(10,2),              -- Total price
    discount_percent INTEGER,                    -- Discount percentage
    spp              INTEGER,                    -- SPP value
    finished_price   NUMERIC(10,2),              -- Final price after discount
    price_with_disc  NUMERIC(10,2),              -- Price with discount
    is_cancel        BOOLEAN,                    -- Is order canceled?
    cancel_date      TIMESTAMP WITH TIME ZONE,   -- Cancelation date
    order_type       TEXT,                       -- Order type
    sticker          TEXT,                       -- Sticker
    g_number         TEXT,                       -- Global number
    srid             TEXT                        -- SRID
);


COMMIT;