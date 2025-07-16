-- Create database
CREATE DATABASE IF NOT EXISTS webcrawler;
USE webcrawler;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME DEFAULT NULL
);

-- Create crawl_results table
CREATE TABLE IF NOT EXISTS crawl_results (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    url TEXT NOT NULL,
    html_version VARCHAR(50),
    title TEXT,
    h1_count INT DEFAULT 0,
    h2_count INT DEFAULT 0,
    h3_count INT DEFAULT 0,
    h4_count INT DEFAULT 0,
    h5_count INT DEFAULT 0,
    h6_count INT DEFAULT 0,
    internal_links INT DEFAULT 0,
    external_links INT DEFAULT 0,
    inaccessible_links INT DEFAULT 0,
    has_login_form BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'queued',
    user_id BIGINT UNSIGNED,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
