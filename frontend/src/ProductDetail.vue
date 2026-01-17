<template>
  <v-app>
    <v-main class="detail-bg">
      <v-container class="py-10 d-flex align-center justify-center" style="min-height: 100vh;">
        <v-card v-if="product" class="pa-12 detail-card" max-width="720" elevation="12" rounded="2xl">
          <v-img :src="product.imageURL || 'https://via.placeholder.com/700x420?text=Resim+Yok'" height="420px" cover class="mb-8 detail-img mx-auto" />
          <div class="detail-price-green text-center mx-auto mb-4">
            <span>{{ product.price }} TL</span>
          </div>
          <h2 class="detail-title mb-4 text-center">{{ product.name }}</h2>
          <div class="detail-desc text-center mx-auto mb-6">
            {{ product.description }}
          </div>
          <div class="stock-info text-center mx-auto mb-6">
            <v-chip
              :color="product.stock > 0 ? 'success' : 'error'"
              class="stock-chip"
              size="large"
            >
              {{ product.stock > 0 ? `${product.stock} in stock` : 'Stokta yok' }}
            </v-chip>
          </div>
          <div class="d-flex justify-center gap-4">
            <v-btn
              v-if="product.stock > 0"
              color="primary"
              @click="addToCart"
              prepend-icon="mdi-cart-plus"
              rounded="xl"
              size="large"
              elevation="4"
              :loading="loading"
              :disabled="loading"
            >
              Add to Cart
            </v-btn>
            <v-btn
              color="info"
              @click="goGallery"
              prepend-icon="mdi-arrow-left"
              rounded="xl"
              size="large"
              elevation="4"
            >
              Go to Gallery
            </v-btn>
          </div>
        </v-card>
        <v-alert v-else type="error" class="mx-auto" max-width="400">Ürün bulunamadı.</v-alert>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { api } from './api/productApi';

const route = useRoute();
const router = useRouter();
const product = ref(null);
const loading = ref(false); // Add loading state

onMounted(async () => {
  try {
    const res = await api.get(`/products/${route.params.id}`);
    product.value = res.data;
    console.log('Product loaded:', product.value); // Debug log
  } catch (error) {
    console.error('Error fetching product:', error);
  }
});

function goGallery() {
  router.push('/gallery');
}

async function addToCart() {
  if (!product.value || product.value.stock <= 0 || loading.value) return;
  
  loading.value = true;
  console.log('Adding to cart:', product.value); // Debug log
  
  try {
    // Fetch latest stock information
    const response = await api.get(`/products/${product.value.id}`);
    const currentStock = response.data.stock;
    console.log('Current stock:', currentStock); // Debug log
    
    if (currentStock <= 0) {
      alert('Üzgünüz, ürün stokta kalmamış.');
      product.value.stock = currentStock; // Update displayed stock
      return;
    }
    
    // Get cart from localStorage
    const cartItems = JSON.parse(localStorage.getItem('cart') || '[]');
    const existingItem = cartItems.find(item => item.id === product.value.id);
    
    if (existingItem) {
      if (existingItem.quantity >= currentStock) {
        alert('Üzgünüz, bu üründen daha fazla stok bulunmamaktadır.');
        return;
      }
      existingItem.quantity++;
      console.log('Updated existing item:', existingItem); // Debug log
    } else {
      cartItems.push({
        ...response.data, // Use latest product data
        quantity: 1
      });
      console.log('Added new item to cart'); // Debug log
    }
    
    // Update cart
    localStorage.setItem('cart', JSON.stringify(cartItems));
    alert('Ürün sepete eklendi!');
    
    // Update displayed stock
    product.value.stock = currentStock;
  } catch (err) {
    console.error('Stok kontrolü yapılırken hata:', err);
    alert('Stok kontrolü yapılamadı. Lütfen daha sonra tekrar deneyin.');
  } finally {
    loading.value = false;
  }
}
</script>

<style>
.detail-bg {
  background: linear-gradient(135deg, #f8ffae 0%, #43c6ac 100%);
  min-height: 100vh;
}
.detail-card {
  background: #fffbe7;
  border: 2px solid #ffd54f;
  box-shadow: 0 12px 48px #43c6ac33, 0 4px 16px #ffd54f33;
  border-radius: 40px;
  padding: 48px !important;
  max-width: 720px;
}
.detail-img {
  border-radius: 32px;
  box-shadow: 0 8px 32px #ffd54f55;
  object-fit: cover;
  width: 100%;
  max-width: 520px;
  height: 420px !important;
  display: block;
  margin-left: auto;
  margin-right: auto;
}
.detail-title {
  font-size: 1.1rem;
  font-weight: 900;
  letter-spacing: 1px;
  color: #43c6ac;
  text-shadow: 0 2px 12px #43c6ac33, 0 1px 0 #fffbe7;
}
.detail-price {
  font-size: 2rem;
  font-weight: bold;
  color: #fff;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  border-radius: 18px;
  padding: 18px 36px;
  box-shadow: 0 2px 12px #ffd54f55;
  margin-right: 0;
  margin-bottom: 0;
  display: inline-block;
}
.detail-desc {
  font-size: 1.3rem;
  color: #333;
  background: #e1f5fe;
  border-radius: 18px;
  padding: 18px 32px;
  box-shadow: 0 2px 12px #81d4fa55;
  display: inline-block;
  min-width: 220px;
  max-width: 100%;
}
.detail-price-green {
  font-size: 1.05rem;
  font-weight: bold;
  color: #fff;
  background: linear-gradient(90deg, #43ea7a 0%, #43c6ac 100%);
  border-radius: 18px;
  padding: 18px 36px;
  box-shadow: 0 2px 12px #43c6ac55;
  margin: 0 auto 24px auto;
  display: block;
  width: fit-content;
  min-width: 180px;
}
.stock-info {
  margin: 24px 0;
}
.stock-chip {
  font-size: 1.1rem !important;
  padding: 24px !important;
  font-weight: 500;
}
@media (max-width: 900px) {
  .detail-card {
    padding: 24px !important;
    max-width: 98vw;
  }
  .detail-img {
    height: 280px !important;
    max-width: 98vw;
  }
}
@media (max-width: 600px) {
  .detail-card {
    padding: 10px !important;
    max-width: 100vw;
  }
  .detail-img {
    height: 160px !important;
    max-width: 100vw;
  }
  .detail-price, .detail-desc {
    font-size: 1rem;
    padding: 10px 12px;
  }
  .detail-price-green {
    font-size: 1rem;
    padding: 10px 16px;
    min-width: 120px;
  }
}
</style> 