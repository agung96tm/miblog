-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "blog_posts" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "title" varchar(200) NOT NULL,
    "body" text NOT NULL,
    "user_id" bigint NOT NULL,
    "created_at" TIMESTAMP NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL
    );

ALTER TABLE "blog_posts" ADD CONSTRAINT "blog_posts_user_id_fk_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "blog_posts_user_id" ON "blog_posts" ("user_id");
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "blog_posts" CASCADE;
COMMIT;
-- +goose StatementEnd