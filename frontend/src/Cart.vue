<template>
  <v-app>
    <v-main class="cart-bg">
      <v-container class="py-10">
        <div class="cart-header text-center mb-8">
          <h1 class="cart-title">Shopping Cart</h1>
          <div class="cart-divider mx-auto my-4"></div>
        </div>

        <!-- Hata mesajı -->
        <v-alert
          v-if="error"
          type="error"
          class="mb-4"
          closable
          @click:close="error = ''"
        >
          {{ error }}
        </v-alert>

        <v-card class="cart-card mx-auto" max-width="700" elevation="10" rounded="xl">
          <v-card-text>
            <div v-if="cart.length === 0" class="text-center py-10">
              <v-icon size="60" color="grey">mdi-cart-off</v-icon>
              <div class="text-h6 mt-2">Your cart is empty.</div>
              <v-btn color="primary" class="mt-4" @click="goGallery" prepend-icon="mdi-arrow-left">Go to Gallery</v-btn>
            </div>
            <div v-else>
              <v-table class="cart-table mb-6">
                <thead>
                  <tr>
                    <th>Product</th>
                    <th class="text-center">Price</th>
                    <th class="text-center">Quantity</th>
                    <th class="text-center">Total</th>
                    <th class="text-center">Remove</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in cart" :key="item.id">
                    <td>
                      <div class="d-flex align-center gap-4">
                        <v-img :src="item.imageUrl || 'https://via.placeholder.com/60x40?text=No+Image'" width="60" height="40" class="rounded" />
                        <span class="font-weight-bold">{{ item.name }}</span>
                      </div>
                    </td>
                    <td class="text-center">{{ item.price }} TL</td>
                    <td class="text-center">
                      <v-btn icon size="small" @click="decrease(item)" :disabled="loading"><v-icon>mdi-minus</v-icon></v-btn>
                      <span class="mx-2">{{ item.quantity }}</span>
                      <v-btn icon size="small" @click="increase(item)" :disabled="loading"><v-icon>mdi-plus</v-icon></v-btn>
                    </td>
                    <td class="text-center">{{ (item.price * item.quantity).toFixed(2) }} TL</td>
                    <td class="text-center">
                      <v-btn icon color="error" @click="remove(item)" :disabled="loading"><v-icon>mdi-delete</v-icon></v-btn>
                    </td>
                  </tr>
                </tbody>
              </v-table>
              <div class="d-flex justify-end align-center mb-4">
                <span class="text-h6 font-weight-bold mr-4">Total:</span>
                <span class="cart-total">{{ totalPrice }} TL</span>
              </div>
              <div class="d-flex justify-end">
                <v-btn
                  color="success"
                  size="large"
                  rounded="xl"
                  @click="checkout"
                  :loading="loading"
                  :disabled="loading"
                  prepend-icon="mdi-credit-card"
                >
                  Complete Order
                </v-btn>
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { api } from './api/productApi';

// Basit global sepet state'i (gerçek uygulamada Vuex/pinia önerilir)
const cart = ref(JSON.parse(localStorage.getItem('cart') || '[]'));
const router = useRouter();
const loading = ref(false);
const error = ref('');

const totalPrice = computed(() =>
  cart.value.reduce((sum, item) => sum + item.price * item.quantity, 0).toFixed(2)
);

function remove(item) {
  cart.value = cart.value.filter(p => p.id !== item.id);
  localStorage.setItem('cart', JSON.stringify(cart.value));
}

async function increase(item) {
  try {
    // Önce ürünün güncel stok bilgisini kontrol et
    const response = await api.get(`/products/${item.id}`);
    const currentStock = response.data.stock;

    if (item.quantity >= currentStock) {
      alert('Üzgünüz, bu üründen daha fazla stok bulunmamaktadır.');
      return;
    }

    item.quantity++;
    localStorage.setItem('cart', JSON.stringify(cart.value));
  } catch (err) {
    console.error('Stok kontrolü yapılırken hata:', err);
    alert('Stok kontrolü yapılamadı. Lütfen daha sonra tekrar deneyin.');
  }
}

function decrease(item) {
  if (item.quantity > 1) {
    item.quantity--;
    localStorage.setItem('cart', JSON.stringify(cart.value));
  }
}

async function checkout() {
  if (loading.value) return;
  loading.value = true;
  error.value = '';

  try {
    // Her ürün için stok güncellemesi yap
    for (const item of cart.value) {
      try {
        await api.put(`/products/${item.id}/stock`, {
          quantity: item.quantity
        });
      } catch (err) {
        if (err.response?.status === 400) {
          error.value = `Üzgünüz, "${item.name}" ürününde yeterli stok kalmamış.`;
        } else {
          error.value = 'Satın alma işlemi sırasında bir hata oluştu.';
        }
        loading.value = false;
        return;
      }
    }

    // Başarılı satın alma
    alert('Siparişiniz başarıyla tamamlandı!');
    cart.value = [];
    localStorage.setItem('cart', '[]');
    router.push('/gallery');
  } catch (err) {
    console.error('Checkout error:', err);
    error.value = 'Satın alma işlemi sırasında bir hata oluştu.';
  } finally {
    loading.value = false;
  }
}

function goGallery() {
  router.push('/gallery');
}
</script>

<style>
.cart-bg {
  background: linear-gradient(135deg, #f8ffae 0%, #43c6ac 100%);
  min-height: 100vh;
}
.cart-header {
  margin-bottom: 32px;
}
.cart-title {
  font-size: 2.2rem;
  font-weight: 900;
  letter-spacing: 2px;
  color: #43c6ac;
  text-shadow: 0 4px 24px #43c6ac33, 0 1px 0 #fffbe7;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
.cart-divider {
  width: 120px;
  height: 6px;
  border-radius: 3px;
  background: linear-gradient(90deg, #ffd54f 0%, #43c6ac 100%);
  box-shadow: 0 2px 12px #ffd54f55;
}
.cart-card {
  background: #fffbe7;
  border: 2px solid #ffd54f;
  box-shadow: 0 8px 32px #43c6ac33, 0 2px 8px #ffd54f33;
  border-radius: 32px;
}
.cart-table th {
  font-size: 1.1rem;
  font-weight: bold;
  background: #e1f5fe;
  color: #43c6ac;
  border-bottom: 2px solid #ffd54f;
}
.cart-table td {
  font-size: 1rem;
  padding: 12px 8px;
}
.cart-total {
  font-size: 1.3rem;
  font-weight: bold;
  color: #43c6ac;
}
</style> 