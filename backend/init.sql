-- Veritabanını oluştur
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'ProductDB')
BEGIN
    CREATE DATABASE ProductDB;
END
GO

USE ProductDB;
GO

-- Users tablosunu oluştur
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Users]') AND type in (N'U'))
BEGIN
    CREATE TABLE Users (
        ID INT IDENTITY(1,1) PRIMARY KEY,
        Name NVARCHAR(100) NOT NULL,
        Email NVARCHAR(100) NOT NULL UNIQUE,
        Password NVARCHAR(100) NOT NULL
    );
END
GO

-- Products tablosunu oluştur
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Products]') AND type in (N'U'))
BEGIN
    CREATE TABLE Products (
        ID INT IDENTITY(1,1) PRIMARY KEY,
        Name NVARCHAR(100) NOT NULL,
        Description NVARCHAR(MAX),
        Price DECIMAL(18,2) NOT NULL,
        Stock INT NOT NULL,
        imageURL NVARCHAR(500)
    );
END
GO

-- Örnek ürünler ekle
IF NOT EXISTS (SELECT * FROM Products)
BEGIN
    INSERT INTO Products (Name, Description, Price, Stock, imageURL)
    VALUES 
    ('Laptop', 'High performance laptop', 999.99, 10, 'https://example.com/laptop.jpg'),
    ('Smartphone', 'Latest model smartphone', 699.99, 15, 'https://example.com/phone.jpg'),
    ('Headphones', 'Wireless noise-canceling headphones', 199.99, 20, 'https://example.com/headphones.jpg');
END
GO 
