<template>
  <v-app>
    <v-main class="product-form-bg">
      <!-- Header with Dark Mode Toggle and Gallery Button -->
      <v-app-bar 
        :color="isDark ? 'grey-darken-4' : 'white'" 
        :elevation="isDark ? 0 : 2"
        class="px-4"
      >
        <v-spacer></v-spacer>
        
        <!-- Gallery Button -->
        <v-btn
          color="primary"
          variant="outlined"
          prepend-icon="mdi-view-gallery"
          @click="goToGallery"
          class="mr-4"
          rounded="xl"
        >
          Gallery
        </v-btn>
        
        <!-- Dark Mode Toggle -->
        <v-btn
          :icon="isDark ? 'mdi-weather-sunny' : 'mdi-weather-night'"
          @click="toggleDarkMode"
          variant="text"
          size="large"
          :color="isDark ? 'yellow' : 'grey-darken-3'"
        ></v-btn>
      </v-app-bar>

      <!-- Stok Uyarı Banner -->
      <v-banner
        v-if="lowStockProducts.length > 0"
        class="stock-warning-banner"
        sticky
        elevation="24"
        rounded="xl"
      >
        <template v-slot:prepend>
          <v-icon
            color="white"
            icon="mdi-alert-decagram"
            size="x-large"
            class="animated-icon"
          ></v-icon>
        </template>

        <v-banner-text class="text-white">
          <div class="banner-header">
            <h2 class="text-h5 font-weight-bold mb-2">🔔 Stock Alert</h2>
            <p class="text-subtitle-1">The following products have critical stock levels:</p>
          </div>
          <v-list density="compact" bg-color="transparent" class="stock-warning-list">
            <v-list-item
              v-for="product in lowStockProducts"
              :key="product.id"
              :title="product.name"
              :subtitle="product.stock === 0 ? 'Out of Stock!' : `Only ${product.stock} left`"
              class="text-white warning-item"
            >
              <template v-slot:prepend>
                <v-icon :color="product.stock === 0 ? 'pink-lighten-3' : 'cyan-lighten-3'" size="large">
                  {{ product.stock === 0 ? 'mdi-package-variant-remove' : 'mdi-package-variant' }}
                </v-icon>
              </template>
            </v-list-item>
          </v-list>
        </v-banner-text>

        <template v-slot:actions>
          <v-btn
            color="white"
            variant="elevated"
            size="large"
            rounded="xl"
            class="update-stock-btn"
            @click="scrollToProductList"
            prepend-icon="mdi-refresh"
          >
            Update Stock
          </v-btn>
        </template>
      </v-banner>

      <!-- Animated Background Shapes -->
      <div class="bg-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
        <div class="shape shape-4"></div>
      </div>
      
      <v-container class="py-8 d-flex align-center justify-center" style="min-height: 100vh;">
        <v-card class="product-form-card mx-auto" max-width="1200" elevation="20" rounded="2xl">
          <!-- Header Section -->
          <div class="form-header pa-8 text-center">
            <v-avatar size="80" class="mb-4" color="primary">
              <v-icon size="48" color="white">
                {{ isEditing ? 'mdi-pencil' : 'mdi-plus-circle' }}
              </v-icon>
            </v-avatar>
            <h1 class="form-title mb-2">
              {{ isEditing ? 'Edit Product' : 'Add New Product' }}
            </h1>
            <p class="form-subtitle">
              {{ isEditing ? 'Update your product information' : 'Create a new product for your store' }}
            </p>
          </div>

          <!-- Form Section -->
          <v-card-text class="pa-8">
            <v-form @submit.prevent="handleSubmit" class="product-form">
              <v-row>
                <!-- Product Name -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="product.name"
                    label="Product Name"
                    prepend-inner-icon="mdi-tag"
                    required
                    :rules="[v => !!v || 'Product name is required']"
                    variant="solo"
                    rounded="xl"
                    class="form-field"
                    color="primary"
                  />
                </v-col>

                <!-- Price -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model.number="product.price"
                    label="Price"
                    type="number"
                    prepend-inner-icon="mdi-currency-usd"
                    required
                    :rules="[
                      v => !!v || 'Price is required',
                      v => v > 0 || 'Price must be greater than 0'
                    ]"
                    variant="solo"
                    rounded="xl"
                    class="form-field"
                    color="primary"
                  />
                </v-col>

                <!-- Stock -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model.number="product.stock"
                    label="Stock Quantity"
                    type="number"
                    prepend-inner-icon="mdi-package-variant"
                    required
                    :rules="[
                      v => !!v || 'Stock is required',
                      v => v >= 0 || 'Stock cannot be negative'
                    ]"
                    variant="solo"
                    rounded="xl"
                    class="form-field"
                    color="primary"
                  />
                </v-col>

                <!-- Image URL -->
                <v-col cols="12" md="6">
                  <v-text-field
                    v-model="product.imageURL"
                    label="Image URL"
                    prepend-inner-icon="mdi-image"
                    variant="solo"
                    rounded="xl"
                    class="form-field"
                    color="primary"
                    placeholder="https://example.com/image.jpg"
                  />
                </v-col>

                <!-- Description -->
                <v-col cols="12">
                  <v-textarea
                    v-model="product.description"
                    label="Product Description"
                    prepend-inner-icon="mdi-text"
                    required
                    :rules="[v => !!v || 'Description is required']"
                    variant="solo"
                    rounded="xl"
                    class="form-field"
                    color="primary"
                    rows="4"
                    auto-grow
                  />
                </v-col>
              </v-row>

              <!-- Action Buttons -->
              <div class="form-actions d-flex justify-center align-center mt-8">
                <v-btn
                  color="primary"
                  type="submit"
                  :loading="loading"
                  size="large"
                  rounded="xl"
                  :prepend-icon="isEditing ? 'mdi-content-save' : 'mdi-plus'"
                  class="action-btn"
                  elevation="4"
                >
                  {{ isEditing ? 'Update Product' : 'Add Product' }}
                </v-btn>
              </div>
            </v-form>
          </v-card-text>
        </v-card>
      </v-container>

      <!-- Products List Section - Only show when not editing -->
      <v-container class="py-8" v-if="!isEditing">
        <v-card class="products-list-card mx-auto" max-width="1200" elevation="20" rounded="2xl">
          <v-card-title class="text-center pa-6">
            <v-icon size="32" color="primary" class="mr-3">mdi-format-list-bulleted</v-icon>
            <span class="text-h4 font-weight-bold">Product Management</span>
          </v-card-title>
          
          <v-card-text class="pa-6">
            <!-- Search and Filter -->
            <v-row class="mb-6">
              <v-col cols="12" md="6">
                <v-text-field
                  v-model="search"
                  label="Search Products"
                  prepend-inner-icon="mdi-magnify"
                  variant="solo"
                  rounded="xl"
                  clearable
                />
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  v-model="sortOrder"
                  :items="sortOptions"
                  label="Sort By"
                  variant="solo"
                  rounded="xl"
                  prepend-inner-icon="mdi-sort"
                />
              </v-col>
            </v-row>

            <!-- Products List - Simple Table Layout -->
            <v-data-table
              :headers="tableHeaders"
              :items="sortedProducts"
              :search="search"
              :items-per-page="10"
              class="elevation-1 rounded-lg"
              :loading="tableLoading"
            >
              <template #item.price="{ item }">
                <v-chip color="primary" size="small" class="font-weight-bold">
                  ${{ item.price }}
                </v-chip>
              </template>
              
              <template #item.stock="{ item }">
                <v-chip 
                  :color="item.stock > 10 ? 'success' : item.stock > 0 ? 'warning' : 'error'" 
                  size="small"
                >
                  {{ item.stock }}
                </v-chip>
              </template>
              
              <template #item.description="{ item }">
                <span class="text-truncate" style="max-width: 200px; display: block;">
                  {{ item.description }}
                </span>
              </template>
              
              <template #item.actions="{ item }">
                <div class="d-flex gap-2">
                  <v-btn
                    color="info"
                    size="small"
                    rounded="xl"
                    prepend-icon="mdi-eye"
                    @click="viewProduct(item.id)"
                  >
                    View
                  </v-btn>
                  <v-btn
                    color="warning"
                    size="small"
                    rounded="xl"
                    prepend-icon="mdi-pencil"
                    @click="editProduct(item)"
                  >
                    Edit
                  </v-btn>
                  <v-btn
                    color="error"
                    size="small"
                    rounded="xl"
                    prepend-icon="mdi-delete"
                    @click="deleteProduct(item.id)"
                  >
                    Delete
                  </v-btn>
                </div>
              </template>
            </v-data-table>
          </v-card-text>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue';
import { createProduct, updateProduct, getProducts, deleteProduct as apiDeleteProduct, getProductById } from '../api/productApi';
import { useRouter, useRoute } from 'vue-router';

const props = defineProps({
  isEditing: {
    type: Boolean,
    default: false
  }
});

const router = useRouter();
const route = useRoute();
const loading = ref(false);
const tableLoading = ref(false);
const isDark = ref(false);
const product = ref({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  imageURL: ''
});

// Products list state
const products = ref([]);
const search = ref('');
const sortOrder = ref('name');

const sortOptions = [
  { title: 'Name A-Z', value: 'name' },
  { title: 'Name Z-A', value: 'name-desc' },
  { title: 'Price Low-High', value: 'price' },
  { title: 'Price High-Low', value: 'price-desc' },
  { title: 'Stock Low-High', value: 'stock' },
  { title: 'Stock High-Low', value: 'stock-desc' }
];

const tableHeaders = [
  { title: 'ID', key: 'id', sortable: true },
  { title: 'Name', key: 'name', sortable: true },
  { title: 'Description', key: 'description', sortable: false },
  { title: 'Price', key: 'price', sortable: true },
  { title: 'Stock', key: 'stock', sortable: true },
  { title: 'Actions', key: 'actions', sortable: false }
];

// Stok uyarı sistemi için yeni değişkenler
const LOW_STOCK_THRESHOLD = 5; // 5 veya daha az stok "düşük stok" olarak kabul edilir
const lowStockProducts = computed(() => {
  return products.value.filter(p => p.stock <= LOW_STOCK_THRESHOLD);
});

// Stok listesine scroll yapma fonksiyonu
function scrollToProductList() {
  const productList = document.querySelector('.products-list-card');
  if (productList) {
    productList.scrollIntoView({ behavior: 'smooth' });
  }
}

// Ürünleri periyodik olarak kontrol et
let stockCheckInterval;

// Fetch products
const fetchProducts = async () => {
  try {
    tableLoading.value = true;
    products.value = await getProducts();
  } catch (error) {
    console.error('Error fetching products:', error);
  } finally {
    tableLoading.value = false;
  }
};

// Fetch single product for editing
const fetchProductForEdit = async (id) => {
  try {
    loading.value = true;
    const fetchedProduct = await getProductById(id);
    product.value = { ...fetchedProduct };
  } catch (error) {
    console.error('Error fetching product for edit:', error);
    alert('Failed to load product for editing');
  } finally {
    loading.value = false;
  }
};

// Computed properties
const sortedProducts = computed(() => {
  const sorted = [...products.value];
  
  switch (sortOrder.value) {
    case 'name':
      return sorted.sort((a, b) => a.name.localeCompare(b.name));
    case 'name-desc':
      return sorted.sort((a, b) => b.name.localeCompare(a.name));
    case 'price':
      return sorted.sort((a, b) => a.price - b.price);
    case 'price-desc':
      return sorted.sort((a, b) => b.price - a.price);
    case 'stock':
      return sorted.sort((a, b) => a.stock - b.stock);
    case 'stock-desc':
      return sorted.sort((a, b) => b.stock - a.stock);
    default:
      return sorted;
  }
});

// Methods
async function handleSubmit() {
  try {
    loading.value = true;
    if (props.isEditing) {
      await updateProduct(product.value.id, product.value);
      alert('Product updated successfully!');
      router.push('/product/new');
    } else {
      await createProduct(product.value);
      // Reset form
      product.value = {
        name: '',
        description: '',
        price: 0,
        stock: 0,
        imageURL: ''
      };
      // Refresh products list
      await fetchProducts();
    }
  } catch (error) {
    console.error('Error saving product:', error);
    alert('Failed to save product. Please try again.');
  } finally {
    loading.value = false;
  }
}

function viewProduct(id) {
  router.push(`/product/${id}`);
}

function editProduct(product) {
  router.push(`/product/${product.id}/edit`);
}

function goToGallery() {
  router.push('/gallery');
}

function toggleDarkMode() {
  isDark.value = !isDark.value;
  // Apply dark mode to v-app
  const app = document.querySelector('.v-application');
  if (app) {
    app.classList.toggle('v-theme--dark', isDark.value);
  }
}

async function deleteProduct(id) {
  if (confirm('Are you sure you want to delete this product?')) {
    try {
      await apiDeleteProduct(id);
      await fetchProducts();
    } catch (error) {
      console.error('Error deleting product:', error);
      alert('Failed to delete product. Please try again.');
    }
  }
}

// Lifecycle
onMounted(async () => {
  if (props.isEditing && route.params.id) {
    await fetchProductForEdit(parseInt(route.params.id));
  } else {
    await fetchProducts();
  }
  
  // İlk ürün listesini yükle
  await fetchProducts();
  
  // Her 5 dakikada bir stok durumunu kontrol et
  stockCheckInterval = setInterval(async () => {
    await fetchProducts();
  }, 5 * 60 * 1000);
});

// Component unmount olduğunda interval'i temizle
onUnmounted(() => {
  if (stockCheckInterval) {
    clearInterval(stockCheckInterval);
  }
});

// Watch for route changes
watch(() => route.params.id, async (newId) => {
  if (props.isEditing && newId) {
    await fetchProductForEdit(parseInt(newId));
  }
});
</script>

<style scoped>
.product-form-bg {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

/* Dark Mode Styles */
.v-theme--dark .product-form-bg {
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
}

.v-theme--dark .product-form-card,
.v-theme--dark .products-list-card {
  background: rgba(30, 30, 30, 0.95);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: white;
}

.v-theme--dark .form-header {
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  color: white;
}

.v-theme--dark .form-title {
  background: linear-gradient(45deg, #ffffff, #e0e0e0);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.v-theme--dark .v-data-table {
  background: rgba(30, 30, 30, 0.95);
  color: white;
}

.v-theme--dark .v-data-table :deep(.v-data-table-header) {
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  color: white;
}

.v-theme--dark .v-data-table :deep(.v-data-table__wrapper) {
  background: rgba(30, 30, 30, 0.95);
}

.v-theme--dark .v-data-table :deep(.v-data-table__tr) {
  background: rgba(40, 40, 40, 0.95);
}

.v-theme--dark .v-data-table :deep(.v-data-table__tr:hover) {
  background: rgba(60, 60, 60, 0.95);
}

.bg-shapes {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 0;
}

.shape {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  animation: float 6s ease-in-out infinite;
  transition: all 0.3s ease;
}

.v-theme--dark .shape {
  background: rgba(255, 255, 255, 0.05);
}

.shape-1 {
  width: 200px;
  height: 200px;
  top: 10%;
  left: 10%;
  animation-delay: 0s;
}

.shape-2 {
  width: 150px;
  height: 150px;
  top: 20%;
  right: 15%;
  animation-delay: 2s;
}

.shape-3 {
  width: 100px;
  height: 100px;
  bottom: 20%;
  left: 20%;
  animation-delay: 4s;
}

.shape-4 {
  width: 120px;
  height: 120px;
  bottom: 10%;
  right: 10%;
  animation-delay: 1s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-20px) rotate(180deg);
  }
}

.product-form-card, .products-list-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  position: relative;
  z-index: 1;
  transition: all 0.3s ease;
}

.form-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border-radius: 24px 24px 0 0;
  transition: all 0.3s ease;
}

.form-title {
  font-size: 2.5rem;
  font-weight: 900;
  background: linear-gradient(45deg, #ffffff, #f0f0f0);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.form-subtitle {
  font-size: 1.1rem;
  opacity: 0.9;
  font-weight: 300;
  transition: all 0.3s ease;
}

.form-field {
  margin-bottom: 16px;
}

.form-field :deep(.v-field) {
  border-radius: 16px;
  transition: all 0.3s ease;
}

.form-field :deep(.v-field:hover) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.v-theme--dark .form-field :deep(.v-field:hover) {
  box-shadow: 0 8px 25px rgba(255, 255, 255, 0.1);
}

.form-field :deep(.v-field--focused) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.3);
}

.v-theme--dark .form-field :deep(.v-field--focused) {
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.5);
}

.form-actions {
  border-top: 1px solid #e0e0e0;
  padding-top: 24px;
  transition: all 0.3s ease;
}

.v-theme--dark .form-actions {
  border-top: 1px solid rgba(255, 255, 255, 0.2);
}

.action-btn {
  min-width: 200px;
  font-weight: 600;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.v-theme--dark .action-btn:hover {
  box-shadow: 0 8px 25px rgba(255, 255, 255, 0.1);
}

/* Data Table Styles */
.v-data-table {
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.v-data-table :deep(.v-data-table-header) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  transition: all 0.3s ease;
}

.v-data-table :deep(.v-data-table-header th) {
  color: white !important;
  font-weight: bold;
}

.v-data-table :deep(.v-data-table__wrapper) {
  border-radius: 16px;
  transition: all 0.3s ease;
}

/* App Bar Styles */
.v-app-bar {
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.v-theme--dark .v-app-bar {
  background: rgba(30, 30, 30, 0.95) !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

/* Responsive Design */
@media (max-width: 768px) {
  .form-title {
    font-size: 2rem;
  }
  
  .form-actions {
    flex-direction: column;
    gap: 16px;
  }
  
  .action-btn {
    width: 100%;
  }
}

.stock-warning-banner {
  margin: 16px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  backdrop-filter: blur(15px);
  box-shadow: 
    0 8px 32px rgba(102, 126, 234, 0.4),
    0 4px 16px rgba(118, 75, 162, 0.3) !important;
}

.stock-warning-banner :deep(.v-banner__content) {
  padding: 24px;
  background: linear-gradient(
    135deg,
    rgba(255, 255, 255, 0.15) 0%,
    rgba(255, 255, 255, 0.05) 100%
  );
}

.banner-header {
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.stock-warning-list {
  margin-top: 16px;
  border-left: 3px solid rgba(255, 255, 255, 0.5);
  padding-left: 24px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 0 16px 16px 0;
  backdrop-filter: blur(10px);
}

.warning-item {
  margin: 8px 0;
  border-radius: 12px;
  transition: all 0.3s ease;
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.1) 0%,
    rgba(255, 255, 255, 0.05) 100%
  );
}

.warning-item:hover {
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.2) 0%,
    rgba(255, 255, 255, 0.1) 100%
  );
  transform: translateX(8px);
}

.stock-warning-list :deep(.v-list-item) {
  min-height: 48px;
  padding: 8px 16px;
}

.stock-warning-list :deep(.v-list-item__prepend) {
  margin-right: 16px;
}

.stock-warning-list :deep(.v-list-item-title) {
  font-size: 1.1rem !important;
  font-weight: 600;
  color: white;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.stock-warning-list :deep(.v-list-item-subtitle) {
  font-size: 0.95rem !important;
  color: rgba(255, 255, 255, 0.9);
  font-weight: 500;
}

.update-stock-btn {
  font-weight: 600;
  letter-spacing: 0.5px;
  padding: 0 32px;
  height: 48px;
  background: linear-gradient(135deg, #81d4fa 0%, #64b5f6 100%) !important;
  color: rgba(0, 0, 0, 0.8) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
}

.update-stock-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
  background: linear-gradient(135deg, #64b5f6 0%, #42a5f5 100%) !important;
}

@keyframes pulse {
  0% {
    transform: scale(1) rotate(0deg);
    opacity: 1;
  }
  50% {
    transform: scale(1.2) rotate(180deg);
    opacity: 0.8;
  }
  100% {
    transform: scale(1) rotate(360deg);
    opacity: 1;
  }
}

.animated-icon {
  animation: pulse 3s infinite;
}

/* Responsive Design */
@media (max-width: 600px) {
  .stock-warning-banner {
    margin: 8px;
  }
  
  .stock-warning-banner :deep(.v-banner__content) {
    padding: 16px;
  }
  
  .banner-header h2 {
    font-size: 1.25rem;
  }
  
  .banner-header p {
    font-size: 1rem;
  }
  
  .stock-warning-list {
    margin-top: 12px;
    padding-left: 16px;
  }
  
  .stock-warning-list :deep(.v-list-item-title) {
    font-size: 1rem !important;
  }
  
  .stock-warning-list :deep(.v-list-item-subtitle) {
    font-size: 0.875rem !important;
  }
  
  .update-stock-btn {
    padding: 0 24px;
    height: 40px;
  }
}
</style> 
