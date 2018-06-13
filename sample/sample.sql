/*** DROP DATABASE ***/
DROP DATABASE IF EXISTS goer;

/*** CREATE DATABASE ***/
CREATE DATABASE goer;

/*** CREATE TABLE ***/
DROP TABLE IF EXISTS goer.user;
CREATE TABLE goer.user (
    user_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
    user_name VARCHAR (64) NOT NULL DEFAULT '' COMMENT 'ユーザー名',
    gender_cd  INTEGER UNSIGNED NOT NULL COMMENT '性別コード',
    pref_cd INTEGER UNSIGNED NOT NULL COMMENT '都道府県コード',

    PRIMARY KEY ( user_id )
) COMMENT='ユーザー';

DROP TABLE IF EXISTS goer.gender;
CREATE TABLE goer.gender (
    gender_cd INTEGER UNSIGNED NOT NULL COMMENT '性別コード',
    gender VARCHAR (64) NOT NULL DEFAULT '' COMMENT '性別',

    PRIMARY KEY ( gender_cd )
) COMMENT='性別';

DROP TABLE IF EXISTS goer.pref;
CREATE TABLE goer.pref (
    pref_cd INTEGER UNSIGNED NOT NULL COMMENT '都道府県コード',
    pref_name VARCHAR (64) NOT NULL DEFAULT '' COMMENT '都道府県名',

    PRIMARY KEY ( pref_cd )
) COMMENT='都道府県';

/*** DROP FOREIGN KEY ***/
/*
ALTER TABLE goer.user DROP FOREIGN KEY fk_user__gender;
ALTER TABLE goer.user DROP FOREIGN KEY fk_user__pref;
*/

/*** CREATE FOREIGN KEY ***/
ALTER TABLE goer.user ADD CONSTRAINT fk_user__gender FOREIGN KEY (gender_cd) REFERENCES gender (gender_cd) ON UPDATE CASCADE ON DELETE NO ACTION;
ALTER TABLE goer.user ADD CONSTRAINT fk_user__pref FOREIGN KEY (pref_cd) REFERENCES pref (pref_cd) ON UPDATE CASCADE ON DELETE NO ACTION;
