package repository //Bu dosyanın repository paketine ait olduğunu belirtir.

import (
	"database/sql" //Go'nun standart SQL veritabanı işlemleri paketi.
	"GoProductManagement/models" //Projenin model tanımlarını içeren paket.
)

//CRUD (Create, Read, Update, Delete) işlemlerini tanımlayan bir arayüz.
type ProductRepository interface { 
	Create(product models.Product) (int, error)  // Yeni ürün oluştur                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  
	GetAll() ([]models.Product, error) // Tüm ürünleri getir
	GetByID(id int) (models.Product, error) // ID'ye göre ürün getir
	Update(product models.Product) error // Ürün güncelle
	Delete(id int) error // Ürün sil
	UpdateStock(id int, quantity int) error // Stok güncelle
}

type productRepository struct { // Veritabanı bağlantısını tutan yapı
	db *sql.DB // SQL veritabanı bağlantı nesnesi
}

func NewProductRepository(db *sql.DB) ProductRepository { // Yeni repository örneği oluşturur
	return &productRepository{db: db} // Veritabanı bağlantısı ile initialize eder
}
// Ürün oluşturma metodu
func (r *productRepository) Create(product models.Product) (int, error) { // SQL Server'a özgü INSERT sorgusu (OUTPUT INSERTED.ID ile eklenen kaydın ID'sini alır)
	query := "INSERT INTO Products (Name, Description, Price, Stock, imageURL) OUTPUT INSERTED.ID VALUES (@p1, @p2, @p3, @p4, @p5)" //OUTPUT INSERTED.ID ile eklenen kaydın ID'si alınır.
	var id int // Oluşturulan ürünün ID'si
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.Stock, product.ImageURL).Scan(&id) // Sorguyu çalıştır ve sonucu id değişkenine atar
	if err != nil {  // Hata kontrolü
		return 0, err // Hata durumunda 0 ve hata döndür
	}         
	return id, nil // ID ve nil (hata yok) döndür
}

func (r *productRepository) GetAll() ([]models.Product, error) { // Tüm ürünleri getirme metodu
	query := "SELECT ID, Name, Description, Price, Stock, imageURL FROM Products" // Tüm ürünleri seçen SQL sorgusu
	rows, err := r.db.Query(query) // Sorguyu çalıştır
	if err != nil { // Hata kontrolü
		return nil, err
	}
	defer rows.Close() // Fonksiyon bitiminde rows'u kapat

	var products []models.Product // Ürün listesi 
	for rows.Next() {
		var p models.Product // Tek bir ürün
		var imageURL sql.NullString // NULL olabilen ImageURL için
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &imageURL)  // Satırdaki değerleri struct'a ata
		if err != nil {
			return nil, err
		}
		if imageURL.Valid { // ImageURL NULL ise boş string ata, değilse değerini al
			p.ImageURL = imageURL.String
		} else {
			p.ImageURL = ""
		}
		products = append(products, p) // Ürünü listeye ekle
	}
	return products, nil // Ürün listesini ve nil (hata yok) döndür 
}

func (r *productRepository) GetByID(id int) (models.Product, error) { // ID'ye göre ürün getirme metodu
	query := "SELECT ID, Name, Description, Price, Stock, imageURL FROM Products WHERE ID = @p1" // Belirli ID'ye sahip ürünü getiren SQL sorgusu
	var p models.Product // Ürün struct'ı
	var imageURL sql.NullString // NULL olabilen ImageURL
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &imageURL) // Sorguyu çalıştır ve sonucu struct'a ata
	if err != nil {
		return models.Product{}, err // Hata durumunda boş ürün ve hata döndür
	}
	if imageURL.Valid { // ImageURL NULL ise boş string ata, değilse değerini al
		p.ImageURL = imageURL.String
	} else {
		p.ImageURL = ""
	}
	return p, nil // Ürün ve nil (hata yok) döndür
}

func (r *productRepository) Update(product models.Product) error { // Ürün güncelleme metodu
	query := "UPDATE Products SET Name=@p1, Description=@p2, Price=@p3, Stock=@p4, imageURL=@p5 WHERE ID=@p6" // Ürün bilgilerini güncelleyen SQL sorgusu
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.ImageURL, product.ID)     // Sorguyu çalıştır (Exec sonuç döndürmez)

	return err  // Hata varsa döndür
}

func (r *productRepository) Delete(id int) error { // Ürün silme metodu
	query := "DELETE FROM Products WHERE ID=@p1" // Belirtilen ID'ye sahip ürünü silen SQL sorgusu
	_, err := r.db.Exec(query, id)  // Sorguyu çalıştır 
	return err // Hata varsa döndür
}

// Stok güncelleme metodu
func (r *productRepository) UpdateStock(id int, quantity int) error {
	query := "UPDATE Products SET Stock = Stock - @p1 WHERE ID = @p2 AND Stock >= @p1"
	result, err := r.db.Exec(query, quantity, id)
	if err != nil {
		return err
	}

	// Etkilenen satır sayısını kontrol et
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Eğer hiç satır etkilenmediyse (stok yetersiz veya ürün bulunamadı)
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}