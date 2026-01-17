<template>
  <v-app>
    <v-main class="gallery-bg">
      <!-- Top Full-Width Rotating Banner -->
      <div class="fullwidth-ad-banner ad-animate" :style="{background: topBanner.bg}" v-if="topBanner">
        <div class="ad-banner-content">
          <img v-if="topBanner.img" :src="topBanner.img" class="ad-banner-img" alt="Ad" />
          <div>
            <div class="ad-banner-title">{{ topBanner.title }}</div>
            <div class="ad-banner-desc">{{ topBanner.desc }}</div>
          </div>
        </div>
      </div>
      
      <v-container class="py-8 gallery-main-content">
        <!-- Floating Action Button -->
        <v-btn
          class="floating-action-btn"
          color="primary"
          icon
          size="large"
          @click="scrollToTop"
          v-show="showScrollTop"
        >
          <v-icon>mdi-arrow-up</v-icon>
        </v-btn>

        <!-- Quick View Modal -->
        <v-dialog v-model="quickViewDialog" max-width="600" persistent>
          <v-card class="quick-view-card">
            <v-card-title class="quick-view-title">
              <span>Quick View</span>
              <v-btn icon @click="quickViewDialog = false">
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </v-card-title>
            <v-card-text v-if="selectedProduct">
              <v-row>
                <v-col cols="12" md="6">
                  <v-img 
                    :src="selectedProduct.imageURL || 'https://via.placeholder.com/300x200?text=Resim+Yok'" 
                    height="250" 
                    cover 
                    class="rounded-lg"
                  />
                </v-col>
                <v-col cols="12" md="6">
                  <h3 class="text-h5 font-weight-bold mb-3">{{ selectedProduct.name }}</h3>
                  <p class="text-body-1 mb-4">{{ selectedProduct.description }}</p>
                  <v-chip color="deep-purple-accent-2" size="large" class="mb-4">
                    {{ selectedProduct.price }} TL
                  </v-chip>
                  <div class="d-flex gap-2">
                    <v-btn color="primary" @click="addToCart(selectedProduct)" prepend-icon="mdi-cart-plus">
                      Add to Cart
                    </v-btn>
                    <v-btn variant="outlined" @click="goProduct(selectedProduct.id)" prepend-icon="mdi-eye">
                      View Details
                    </v-btn>
                  </div>
                </v-col>
              </v-row>
            </v-card-text>
          </v-card>
        </v-dialog>

        <div class="action-bar-below-ad d-flex justify-end align-center" style="gap: 10px; margin-bottom: 12px;">
          <v-btn icon class="cart-btn" @click="goCart" title="Cart">
            <v-badge :content="cart.length" color="red" overlap>
              <v-icon size="32">mdi-cart</v-icon>
            </v-badge>
          </v-btn>
          <v-btn icon class="user-btn" @click="goProfile" title="User Profile">
            <v-icon size="32" color="primary">mdi-account-circle</v-icon>
          </v-btn>
          <v-btn icon class="darkmode-btn" @click="toggleDark" :title="isDark ? 'Light Mode' : 'Dark Mode'">
            <v-icon size="32">{{ isDark ? 'mdi-white-balance-sunny' : 'mdi-weather-night' }}</v-icon>
          </v-btn>
        </div>
        
        <!-- Modern Category Filter -->
        <div class="mega-category-bar d-flex justify-center align-center mb-6">
          <v-menu
            v-model="megaMenuOpen"
            :close-on-content-click="false"
            offset-y
            max-width="900"
            min-width="400"
            transition="slide-y-transition"
          >
            <template v-slot:activator="{ props }">
              <v-btn
                v-bind="props"
                class="mega-category-btn"
                color="primary"
                variant="elevated"
                prepend-icon="mdi-menu"
                size="x-large"
                rounded="xl"
                elevation="8"
                height="60"
                min-width="220"
              >
                <span class="text-h6 font-weight-bold">Categories</span>
              </v-btn>
            </template>
            <v-card class="mega-menu-dropdown" elevation="16" rounded="xl">
              <v-row no-gutters>
                <v-col cols="5" class="mega-main-list">
                  <v-list>
                    <v-list-item
                      v-for="cat in megaCategories"
                      :key="cat.name"
                      @mouseenter="selectedMainCategory = cat"
                      @click="selectedMainCategory = cat"
                      :class="{ 'active-main': selectedMainCategory && selectedMainCategory.name === cat.name }"
                      class="mega-main-item"
                      rounded="lg"
                    >
                      <template v-slot:prepend>
                        <v-avatar :color="cat.color" size="40" class="mega-main-avatar">
                          <v-icon color="white" size="22">{{ cat.icon }}</v-icon>
                        </v-avatar>
                      </template>
                      <v-list-item-title class="mega-main-title">{{ cat.name }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-col>
                <v-col cols="7" class="mega-sub-list" v-if="selectedMainCategory">
                  <v-list>
                    <v-list-item
                      v-for="sub in selectedMainCategory.sub"
                      :key="sub.name"
                      @click="selectCategory(sub.name)"
                      class="mega-sub-item"
                      rounded="lg"
                    >
                      <template v-slot:prepend>
                        <v-avatar color="primary" size="32" class="mega-sub-avatar">
                          <v-icon color="white" size="18">{{ sub.icon }}</v-icon>
                        </v-avatar>
                      </template>
                      <v-list-item-title class="mega-sub-title">{{ sub.name }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-col>
              </v-row>
            </v-card>
          </v-menu>
        </div>
        
        <div class="gallery-header text-center mb-6">
          <h1 class="gallery-title modern-gallery-title">Product Gallery</h1>
          <div class="gallery-divider mx-auto my-4"></div>
        </div>
        
        <v-row class="mb-6 align-center">
          <v-col cols="12" md="9" class="d-flex justify-center">
            <v-text-field
              v-model="search"
              label="Search Product"
              prepend-inner-icon="mdi-magnify"
              clearable
              rounded
              variant="solo"
              style="max-width: 800px; width: 100%; background: #fffbe7; border-radius: 32px; margin: 0 auto; box-shadow: 0 2px 12px #ffe08255;"
              color="amber-darken-2"
              class="search-colored"
            />
          </v-col>
          <v-col cols="12" md="3" class="d-flex flex-column align-end">
            <div style="width: 100%; max-width: 800px;">
              <v-select
                v-model="sortOrder"
                :items="sortOptions"
                label="Sort by Price"
                prepend-inner-icon="mdi-sort"
                dense
                hide-details
                variant="solo"
                style="width: 100%; min-width: 200px; background: #e1f5fe; border-radius: 32px; box-shadow: 0 2px 12px #81d4fa55;"
                color="cyan-darken-2"
              />
            </div>
          </v-col>
        </v-row>
        
        <!-- View Toggle Buttons -->
        <div class="view-toggle-container d-flex justify-center mb-6">
          <v-btn-group>
            <v-btn
              :variant="viewMode === 'grid' ? 'elevated' : 'outlined'"
              @click="viewMode = 'grid'"
              prepend-icon="mdi-view-grid"
            >
              Grid
            </v-btn>
            <v-btn
              :variant="viewMode === 'list' ? 'elevated' : 'outlined'"
              @click="viewMode = 'list'"
              prepend-icon="mdi-view-list"
            >
              List
            </v-btn>
          </v-btn-group>
        </div>
        
        <!-- Grid View -->
        <v-row v-if="viewMode === 'grid'">
          <v-col v-for="product in paginatedProducts" :key="product.id" cols="12" sm="6" md="4" lg="3">
            <v-card class="mx-auto gallery-card vibrant-card animated-card" max-width="320" elevation="10" rounded="xl">
              <div class="card-image-container">
                <v-img 
                  :src="product.imageURL || 'https://via.placeholder.com/300x200?text=Resim+Yok'" 
                  height="200px" 
                  cover 
                  class="rounded-t-xl card-image"
                />
                <div class="card-overlay">
                  <v-btn icon class="quick-view-btn" @click="openQuickView(product)">
                    <v-icon>mdi-eye</v-icon>
                  </v-btn>
                  <v-btn icon class="favorite-btn" @click="toggleFavorite(product)">
                    <v-icon :color="isFavorite(product.id) ? 'red' : 'white'">
                      {{ isFavorite(product.id) ? 'mdi-heart' : 'mdi-heart-outline' }}
                    </v-icon>
                  </v-btn>
                </div>
              </div>
              <v-card-title class="font-weight-bold text-center">{{ product.name }}</v-card-title>
              <v-card-actions class="justify-center">
                <v-chip color="deep-purple-accent-2" text-color="white" size="large" class="font-weight-bold price-chip">
                  {{ product.price }} TL
                </v-chip>
              </v-card-actions>
              <v-card-actions class="justify-center pb-4" style="gap: 8px;">
                <v-btn class="go-btn" @click="goProduct(product.id)" rounded="xl" prepend-icon="mdi-eye">
                  View Product
                </v-btn>
                <v-btn class="add-cart-btn" color="primary" @click="addToCart(product)" rounded="xl" prepend-icon="mdi-cart-plus">
                  Add to Cart
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
        
        <!-- List View -->
        <div v-else class="list-view-container">
          <v-card 
            v-for="product in paginatedProducts" 
            :key="product.id" 
            class="list-view-card mb-4"
            elevation="8"
            rounded="xl"
          >
            <v-row align="center">
              <v-col cols="12" md="3">
                <v-img 
                  :src="product.imageURL || 'https://via.placeholder.com/300x200?text=Resim+Yok'" 
                  height="150" 
                  cover 
                  class="rounded-lg"
                />
              </v-col>
              <v-col cols="12" md="6">
                <h3 class="text-h6 font-weight-bold mb-2">{{ product.name }}</h3>
                <p class="text-body-2 text-truncate">{{ product.description }}</p>
                <v-chip color="deep-purple-accent-2" class="mt-2">
                  {{ product.price }} TL
                </v-chip>
              </v-col>
              <v-col cols="12" md="3" class="d-flex flex-column gap-2">
                <v-btn color="primary" @click="addToCart(product)" prepend-icon="mdi-cart-plus" size="small">
                  Add to Cart
                </v-btn>
                <v-btn variant="outlined" @click="goProduct(product.id)" prepend-icon="mdi-eye" size="small">
                  View Details
                </v-btn>
                <v-btn variant="text" @click="openQuickView(product)" prepend-icon="mdi-eye" size="small">
                  Quick View
                </v-btn>
              </v-col>
            </v-row>
          </v-card>
        </div>
        
        <div class="d-flex justify-center my-8">
          <v-pagination
            v-model="currentPage"
            :length="pageCount"
            color="primary"
            size="large"
            total-visible="7"
            @update:modelValue="onPageChange"
          />
        </div>
        
        <div class="d-flex justify-center mt-10 mb-4" style="gap: 16px;">
          <v-btn class="logout-btn" @click="logout" prepend-icon="mdi-logout" size="large" rounded="xl" elevation="4">
            Logout
          </v-btn>
        </div>
      </v-container>
      
      <!-- Bottom Full-Width Rotating Banner -->
      <div class="fullwidth-ad-banner ad-animate" :style="{background: bottomBanner.bg}" v-if="bottomBanner">
        <div class="ad-banner-content">
          <img v-if="bottomBanner.img" :src="bottomBanner.img" class="ad-banner-img" alt="Ad" />
          <div>
            <div class="ad-banner-title">{{ bottomBanner.title }}</div>
            <div class="ad-banner-desc">{{ bottomBanner.desc }}</div>
          </div>
        </div>
      </div>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { getProducts } from './api/productApi';

const cart = ref([]);
const products = ref([]);
const search = ref('');
const sortOrder = ref('artan');
const viewMode = ref('grid'); // 'grid' or 'list'
const sortOptions = [
  { title: 'Price: Low to High', value: 'artan' },
  { title: 'Price: High to Low', value: 'azalan' }
];
const router = useRouter();
const isDark = ref(false);
function toggleDark() {
  isDark.value = !isDark.value;
}

const categories = [
  'All',
  'Electronics',
  'Clothing',
  'Books',
  'Home',
  'Other',
  'Sports',
  'Toys',
  'Beauty',
  'Accessories',
  'Shoes',
  'Jewelry',
  'Music',
  'Art',
  'Office',
  'Garden',
  'Pets',
  'Kids',
  'Automotive',
  'Health',
  'Outdoors',
  'Grocery',
  'Stationery',
  'Games',
  'Watches',
  'Bags',
  'Baby',
  'Furniture',
  'Lighting',
  'Travel',
  'Crafts'
];

const categoryColors = [
  'linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%)',
  'linear-gradient(90deg, #ff6e7f 0%, #bfe9ff 100%)',
  'linear-gradient(90deg, #a1c4fd 0%, #c2e9fb 100%)',
  'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)',
  'linear-gradient(90deg, #f9d423 0%, #ff4e50 100%)',
  'linear-gradient(90deg, #43cea2 0%, #185a9d 100%)',
  'linear-gradient(90deg, #ffaf7b 0%, #d76d77 100%)',
  'linear-gradient(90deg, #a8edea 0%, #fed6e3 100%)',
  'linear-gradient(90deg, #fcb69f 0%, #ffecd2 100%)',
  'linear-gradient(90deg, #a1ffce 0%, #faffd1 100%)',
  'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)',
  'linear-gradient(90deg, #f9d423 0%, #ff4e50 100%)',
  'linear-gradient(90deg, #43cea2 0%, #185a9d 100%)',
  'linear-gradient(90deg, #ffaf7b 0%, #d76d77 100%)',
  'linear-gradient(90deg, #a8edea 0%, #fed6e3 100%)',
  'linear-gradient(90deg, #fcb69f 0%, #ffecd2 100%)',
  'linear-gradient(90deg, #a1ffce 0%, #faffd1 100%)',
  'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)',
  'linear-gradient(90deg, #f9d423 0%, #ff4e50 100%)',
  'linear-gradient(90deg, #43cea2 0%, #185a9d 100%)',
  'linear-gradient(90deg, #ffaf7b 0%, #d76d77 100%)',
  'linear-gradient(90deg, #a8edea 0%, #fed6e3 100%)',
  'linear-gradient(90deg, #fcb69f 0%, #ffecd2 100%)',
  'linear-gradient(90deg, #a1ffce 0%, #faffd1 100%)',
  'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)',
  'linear-gradient(90deg, #f9d423 0%, #ff4e50 100%)',
  'linear-gradient(90deg, #43cea2 0%, #185a9d 100%)',
  'linear-gradient(90deg, #ffaf7b 0%, #d76d77 100%)',
  'linear-gradient(90deg, #a8edea 0%, #fed6e3 100%)'
];

function getCategoryColor(index) {
  return categoryColors[index % categoryColors.length];
}
const selectedCategory = ref('All');
const categoryMenuOpen = ref(false);

// Modern category filter state



function getCategoryIcon(category) {
  const icons = {
    'All': 'mdi-view-grid',
    'Electronics': 'mdi-laptop',
    'Clothing': 'mdi-tshirt-crew',
    'Books': 'mdi-book-open-variant',
    'Home': 'mdi-home',
    'Other': 'mdi-dots-horizontal',
    'Sports': 'mdi-basketball',
    'Toys': 'mdi-toy-brick',
    'Beauty': 'mdi-lipstick',
    'Accessories': 'mdi-watch',
    'Shoes': 'mdi-shoe-heel',
    'Jewelry': 'mdi-diamond-stone',
    'Music': 'mdi-music',
    'Art': 'mdi-palette',
    'Office': 'mdi-briefcase',
    'Garden': 'mdi-flower',
    'Pets': 'mdi-paw',
    'Kids': 'mdi-baby-face',
    'Automotive': 'mdi-car',
    'Health': 'mdi-heart-pulse',
    'Outdoors': 'mdi-tent',
    'Grocery': 'mdi-food-apple',
    'Stationery': 'mdi-pencil',
    'Games': 'mdi-gamepad-variant',
    'Watches': 'mdi-watch',
    'Bags': 'mdi-bag-personal',
    'Baby': 'mdi-baby-bottle',
    'Furniture': 'mdi-sofa',
    'Lighting': 'mdi-lightbulb',
    'Travel': 'mdi-airplane',
    'Crafts': 'mdi-scissors-cutting'
  };
  return icons[category] || 'mdi-tag';
}

function getCategoryCount(category) {
  return products.value.filter(p => (p.category || 'Other') === category).length;
}

// Pagination
const itemsPerPage = 12;
const currentPage = ref(1);
const pageCount = computed(() => Math.ceil(sortedProducts.value.length / itemsPerPage));

const adMessages = [
  'Big Sale! Up to 50% off on Electronics. Limited Time Offer!',
  'Free shipping for orders over $100!',
  'New arrivals: Check out the latest products!',
  'Sign up now and get 10% off your first order!',
  'Hurry! Limited stock available on selected items.'
];
const currentAdIndex = ref(0);
const currentAd = computed(() => adMessages[currentAdIndex.value]);
let adInterval = null;

onMounted(() => {
  adInterval = setInterval(() => {
    currentAdIndex.value = (currentAdIndex.value + 1) % adMessages.length;
  }, 3000);
});
onUnmounted(() => {
  if (adInterval) clearInterval(adInterval);
});

const fetchProducts = async () => {
  try {
    products.value = await getProducts();
  } catch (error) {
    console.error('Products yüklenirken hata:', error);
  }
};

// Component mount olduğunda ve route değiştiğinde ürünleri yükle
onMounted(() => {
  fetchProducts();
});

const filteredProducts = computed(() => {
  let arr = products.value;
  if (selectedCategory.value !== 'All') {
    arr = arr.filter(p => (p.category || 'Other') === selectedCategory.value);
  }
  if (!search.value) return arr;
  return arr.filter(p =>
    p.name.toLowerCase().includes(search.value.toLowerCase())
  );
});

const sortedProducts = computed(() => {
  const arr = [...filteredProducts.value];
  if (sortOrder.value === 'artan') {
    return arr.sort((a, b) => a.price - b.price);
  } else {
    return arr.sort((a, b) => b.price - a.price);
  }
});

const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  return sortedProducts.value.slice(start, start + itemsPerPage);
});

function addToCart(product) {
  const exists = cart.value.find(p => p.id === product.id);
  if (!exists) {
    cart.value.push({ ...product, quantity: 1 });
  } else {
    exists.quantity += 1;
  }
  localStorage.setItem('cart', JSON.stringify(cart.value));
}

function goCart() {
  router.push('/cart');
}

function logout() {
  localStorage.removeItem('isLoggedIn');
  router.push('/login');
}

function goProduct(id) {
  router.push(`/product/${id}`);
}

function goHome() {
  router.push('/');
}

function onPageChange(page) {
  currentPage.value = page;
}

function goProfile() {
  router.push('/profile');
}

// Rotating full-width ad banners (ENGLISH)
const topAdBanners = [
  {
    title: 'Mega Summer Sale!',
    desc: 'Up to 50% OFF on all electronics. Don\'t miss out!',
    img: 'https://images.unsplash.com/photo-1512436991641-6745cdb1723f?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #ffecd2 0%, #fcb69f 100%)'
  },
  {
    title: 'Free Shipping!',
    desc: 'Enjoy free shipping on orders over $100. Shop now!',
    img: 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #a1c4fd 0%, #c2e9fb 100%)'
  },
  {
    title: 'New Arrivals',
    desc: 'Discover the latest trends in our new collection.',
    img: 'https://images.unsplash.com/photo-1515378791036-0648a3ef77b2?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)'
  }
];
const bottomAdBanners = [
  {
    title: 'Sign Up & Save',
    desc: 'Register now and get 10% off your first order!',
    img: 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #f9d423 0%, #ff4e50 100%)'
  },
  {
    title: 'Limited Stock!',
    desc: 'Hurry up! Selected items are running out fast.',
    img: 'https://images.unsplash.com/photo-1461749280684-dccba630e2f6?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #43cea2 0%, #185a9d 100%)'
  },
  {
    title: 'Gift Cards Available',
    desc: 'Give the gift of choice with our store gift cards.',
    img: 'https://images.unsplash.com/photo-1519125323398-675f0ddb6308?auto=format&fit=crop&w=400&q=80',
    bg: 'linear-gradient(90deg, #ffaf7b 0%, #d76d77 100%)'
  }
];
const topBannerIndex = ref(0);
const bottomBannerIndex = ref(0);
const topBanner = computed(() => topAdBanners[topBannerIndex.value]);
const bottomBanner = computed(() => bottomAdBanners[bottomBannerIndex.value]);
let bannerInterval = null;
onMounted(() => {
  bannerInterval = setInterval(() => {
    topBannerIndex.value = (topBannerIndex.value + 1) % topAdBanners.length;
    bottomBannerIndex.value = (bottomBannerIndex.value + 1) % bottomAdBanners.length;
  }, 3000);
});
onUnmounted(() => {
  if (bannerInterval) clearInterval(bannerInterval);
});

// Quick View Modal
const quickViewDialog = ref(false);
const selectedProduct = ref(null);

function openQuickView(product) {
  selectedProduct.value = product;
  quickViewDialog.value = true;
}

// Favorites
const favorites = ref(JSON.parse(localStorage.getItem('favorites') || '[]'));

function isFavorite(productId) {
  return favorites.value.includes(productId);
}

function toggleFavorite(product) {
  const index = favorites.value.indexOf(product.id);
  if (index > -1) {
    favorites.value.splice(index, 1);
  } else {
    favorites.value.push(product.id);
  }
  localStorage.setItem('favorites', JSON.stringify(favorites.value));
}

// Scroll to top
const showScrollTop = ref(false);

onMounted(() => {
  window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});

function handleScroll() {
  showScrollTop.value = window.scrollY > 300;
}

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}

// Kategori veri yapısını güncelle
const megaCategories = [
  {
    name: 'Elektronik',
    icon: 'mdi-laptop',
    color: 'linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%)',
    sub: [
      { name: 'Bilgisayar', icon: 'mdi-monitor' },
      { name: 'Telefon', icon: 'mdi-cellphone' },
      { name: 'TV & Ses', icon: 'mdi-television' },
      { name: 'Aksesuar', icon: 'mdi-headphones' }
    ]
  },
  {
    name: 'Moda',
    icon: 'mdi-tshirt-crew',
    color: 'linear-gradient(90deg, #ff6e7f 0%, #bfe9ff 100%)',
    sub: [
      { name: 'Kadın', icon: 'mdi-gender-female' },
      { name: 'Erkek', icon: 'mdi-gender-male' },
      { name: 'Çocuk', icon: 'mdi-baby-face' }
    ]
  },
  {
    name: 'Ev & Yaşam',
    icon: 'mdi-sofa',
    color: 'linear-gradient(90deg, #a1c4fd 0%, #c2e9fb 100%)',
    sub: [
      { name: 'Mobilya', icon: 'mdi-sofa' },
      { name: 'Dekorasyon', icon: 'mdi-lamp' },
      { name: 'Mutfak', icon: 'mdi-silverware-fork-knife' }
    ]
  },
  {
    name: 'Spor & Outdoor',
    icon: 'mdi-basketball',
    color: 'linear-gradient(90deg, #fbc2eb 0%, #a6c1ee 100%)',
    sub: [
      { name: 'Spor Giyim', icon: 'mdi-tshirt-crew' },
      { name: 'Kamp & Outdoor', icon: 'mdi-tent' }
    ]
  }
];
const selectedMainCategory = ref(null);
const selectedSubCategory = ref(null);
const megaMenuOpen = ref(false);

// Modern category filter state

function selectCategory(cat) {
  selectedCategory.value = cat;
  categoryMenuOpen.value = false;
}

// Pagination
</script>

<style>
.gallery-bg {
  background: v-bind('isDark ? "linear-gradient(135deg, #18122B 0%, #392F5A 60%, #635985 100%)" : "linear-gradient(135deg, #f8ffae 0%, #43c6ac 100%)"');
  min-height: 100vh;
  transition: background 0.5s;
}
.vibrant-card {
  background: v-bind('isDark ? "linear-gradient(135deg, #231942 60%, #39375B 100%)" : "linear-gradient(135deg, #fffbe7 60%, #e1f5fe 100%)"');
  border: 2px solid v-bind('isDark ? "#635985" : "#ffd54f"');
  color: v-bind('isDark ? "#fff" : "#333"');
  box-shadow: 0 8px 32px #18122b55, 0 2px 8px #39375b33;
  transition: background 0.5s, color 0.5s, border 0.5s;
}
.vibrant-card:hover {
  box-shadow: 0 12px 40px #43c6ac55, 0 4px 16px #ffd54f55;
  transform: translateY(-6px) scale(1.04);
}
.price-chip {
  font-size: 1.1rem;
  letter-spacing: 1px;
  box-shadow: 0 2px 8px #7c4dff33;
}
.go-btn {
  background: linear-gradient(90deg, #43c6ac 0%, #f8ffae 100%);
  color: #333 !important;
  font-weight: bold;
  box-shadow: 0 2px 8px #43c6ac33;
  border: none;
  transition: background 0.2s, color 0.2s;
}
.go-btn:hover {
  background: linear-gradient(90deg, #f8ffae 0%, #43c6ac 100%);
  color: #222 !important;
}
.back-btn {
  background: linear-gradient(90deg, #ffd54f 0%, #43c6ac 100%);
  color: #333 !important;
  font-weight: bold;
  box-shadow: 0 2px 8px #ffd54f33;
  border: none;
  transition: background 0.2s, color 0.2s;
}
.back-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  color: #222 !important;
}
.logout-btn {
  background: linear-gradient(90deg, #ff6e7f 0%, #bfe9ff 100%);
  color: #fff !important;
  font-weight: bold;
  box-shadow: 0 2px 8px #ff6e7f33;
  border: none;
  transition: background 0.2s, color 0.2s;
}
.logout-btn:hover {
  background: linear-gradient(90deg, #bfe9ff 0%, #ff6e7f 100%);
  color: #fff !important;
}
.search-colored .v-field__input {
  background: #fffbe7 !important;
  border-radius: 32px !important;
}
.search-colored .v-field__outline {
  border-color: #ffd54f !important;
}
.search-colored .v-icon {
  color: #ffd54f !important;
}
.gallery-header {
  margin-bottom: 32px;
}
.gallery-title {
  font-size: 2.8rem;
  font-weight: 900;
  letter-spacing: 2px;
  color: #43c6ac;
  text-shadow: 0 4px 24px #43c6ac33, 0 1px 0 #fffbe7;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
.gallery-divider {
  width: 120px;
  height: 6px;
  border-radius: 3px;
  background: linear-gradient(90deg, #ffd54f 0%, #43c6ac 100%);
  box-shadow: 0 2px 12px #ffd54f55;
}
.cart-btn {
  margin-right: 8px;
}
.add-cart-btn {
  font-weight: bold;
  box-shadow: 0 2px 8px #43c6ac33;
  transition: background 0.2s, color 0.2s;
}
.add-cart-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  color: #222 !important;
}
.category-bar {
  gap: 8px;
  flex-wrap: wrap;
}
.category-btn {
  min-width: 120px;
  font-weight: bold;
  letter-spacing: 1px;
  box-shadow: 0 2px 8px #43c6ac22;
  transition: background 0.2s, color 0.2s;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%) !important;
  color: #fff !important;
  border: none;
}
.category-btn.primary {
  color: #fff !important;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%) !important;
  box-shadow: 0 4px 16px #43c6ac44;
}
.category-btn:not(.primary) {
  background: #e1f5fe !important;
  color: #43c6ac !important;
  border: 1.5px solid #43c6ac22;
}
.ad-banner {
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(90deg, #ff6e7f 0%, #bfe9ff 100%);
  color: #fff;
  font-size: 1.5rem;
  font-weight: bold;
  border-radius: 32px;
  box-shadow: 0 8px 32px #ff6e7f55, 0 2px 8px #bfe9ff55;
  padding: 28px 48px;
  max-width: 800px;
  border: 3px solid #fffbe7;
  position: relative;
  overflow: hidden;
}
.ad-banner::after {
  content: '';
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: linear-gradient(120deg, #fffbe7 0%, #ffd54f44 100%);
  opacity: 0.13;
  pointer-events: none;
}
.ad-icon-bounce {
  margin-right: 18px;
  animation: bounce 1.2s infinite alternate;
  display: flex;
  align-items: center;
}
@keyframes bounce {
  0% { transform: translateY(0); }
  100% { transform: translateY(-12px) scale(1.12); }
}
.ad-text-strong {
  letter-spacing: 1.5px;
  font-size: 1.5rem;
  color: #fff;
  text-shadow: 0 2px 12px #ff6e7f55, 0 1px 0 #fffbe7;
}
.user-btn {
  margin-left: 4px;
}
.darkmode-btn {
  margin-left: 4px;
}
body[data-theme='dark'] .v-application {
  background: #2a1746 !important;
}
.gallery-flex-container {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  min-height: 100vh;
}
.side-ad-panel {
  width: 140px;
  min-width: 120px;
  max-width: 160px;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  padding-top: 32px;
  height: 100%;
}
.ad-content {
  width: 100%;
  text-align: center;
  font-weight: bold;
  color: #43c6ac;
  margin-top: 12px;
}
.gallery-main-content {
  flex: 1 1 0%;
  min-width: 0;
  max-width: 1200px;
}
.ad-title {
  font-size: 1.1rem;
  font-weight: bold;
  color: #ff6e7f;
  margin-bottom: 8px;
}
.ad-desc {
  font-size: 0.98rem;
  color: #39375B;
  margin-top: 4px;
  font-weight: 500;
}
.fullwidth-ad-banner {
  width: 100vw;
  min-height: 110px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 32px #18122b33, 0 2px 8px #39375b22;
  border-radius: 0 0 32px 32px;
  margin-bottom: 32px;
  margin-top: 0;
  animation: fadeInAd 0.7s;
  position: relative;
  overflow: hidden;
  z-index: 2;
}
.fullwidth-ad-banner:last-of-type {
  border-radius: 32px 32px 0 0;
  margin-top: 32px;
  margin-bottom: 0;
}
.ad-banner-content {
  display: flex;
  align-items: center;
  gap: 32px;
  padding: 24px 48px;
}
.ad-banner-img {
  width: 100px;
  height: 100px;
  object-fit: cover;
  border-radius: 24px;
  box-shadow: 0 2px 12px #43c6ac33;
  background: #fff;
}
.ad-banner-title {
  font-size: 2rem;
  font-weight: 900;
  color: #fff;
  text-shadow: 0 2px 12px #00000033, 0 1px 0 #fffbe7;
  margin-bottom: 8px;
  letter-spacing: 1.5px;
}
.ad-banner-desc {
  font-size: 1.2rem;
  color: #fffbe7;
  font-weight: 600;
  text-shadow: 0 1px 8px #00000022;
}
@media (max-width: 900px) {
  .ad-banner-content {
    flex-direction: column;
    gap: 12px;
    padding: 16px 8px;
  }
  .ad-banner-title {
    font-size: 1.2rem;
  }
  .ad-banner-desc {
    font-size: 1rem;
  }
  .ad-banner-img {
    width: 60px;
    height: 60px;
  }
}
@keyframes fadeInAd {
  from { opacity: 0; transform: translateY(-20px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (max-width: 960px) {
  .side-ad-panel {
    display: none !important;
  }
  .gallery-main-content {
    max-width: 100vw;
  }
}
.vibrant-category-btn {
  border-radius: 32px !important;
  font-weight: bold;
  letter-spacing: 1.5px;
  margin-bottom: 8px;
  margin-top: 8px;
  box-shadow: 0 2px 8px #43c6ac22;
  transition: all 0.25s cubic-bezier(.4,2,.6,1);
  border: 1.5px solid #fffbe7;
  min-width: 120px;
  font-size: 1.08rem;
  padding: 12px 28px;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%) !important;
  color: #fff !important;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.vibrant-category-btn:hover {
  filter: brightness(1.08) saturate(1.2);
  box-shadow: 0 6px 24px #ffd54f55, 0 2px 8px #43c6ac33;
  transform: scale(1.07) rotate(-1deg);
  z-index: 2;
}
.modern-gallery-title {
  font-size: 3.2rem;
  font-weight: 900;
  letter-spacing: 2.5px;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-fill-color: transparent;
  text-shadow: 0 6px 32px #ffd54f33, 0 2px 8px #43c6ac33;
  animation: popTitle 1.2s cubic-bezier(.4,2,.6,1);
}
@keyframes popTitle {
  0% { opacity: 0; transform: scale(0.8) translateY(-30px); }
  80% { opacity: 1; transform: scale(1.08) translateY(6px); }
  100% { opacity: 1; transform: scale(1) translateY(0); }
}
@media (max-width: 700px) {
  .modern-gallery-title {
    font-size: 1.5rem;
    letter-spacing: 1px;
  }
}
.action-bar-below-ad {
  margin-top: 0;
  margin-bottom: 12px;
  width: 100%;
  max-width: 1200px;
  padding-right: 0;
}
@media (max-width: 900px) {
  .action-bar-below-ad {
    justify-content: flex-end !important;
    margin-bottom: 8px;
    max-width: 100vw;
    padding-right: 0;
  }
}
.view-toggle-container {
  margin-bottom: 24px;
}
.view-toggle-container .v-btn-group {
  background: #e1f5fe;
  border-radius: 32px;
  box-shadow: 0 2px 12px #81d4fa55;
  padding: 4px;
}
.view-toggle-container .v-btn-group .v-btn {
  border-radius: 28px;
  font-weight: bold;
  letter-spacing: 1px;
  box-shadow: none;
  transition: background 0.2s, color 0.2s;
}
.view-toggle-container .v-btn-group .v-btn:hover {
  background: #43c6ac;
  color: #fff;
}
.view-toggle-container .v-btn-group .v-btn.v-btn--elevated {
  background: #43c6ac;
  color: #fff;
}
.view-toggle-container .v-btn-group .v-btn.v-btn--outlined {
  background: #e1f5fe;
  color: #43c6ac;
}
.view-toggle-container .v-btn-group .v-btn.v-btn--outlined:hover {
  background: #43c6ac;
  color: #fff;
}
.list-view-container .v-card {
  background: v-bind('isDark ? "linear-gradient(135deg, #231942 60%, #39375B 100%)" : "linear-gradient(135deg, #fffbe7 60%, #e1f5fe 100%)"');
  border: 2px solid v-bind('isDark ? "#635985" : "#ffd54f"');
  color: v-bind('isDark ? "#fff" : "#333"');
  box-shadow: 0 8px 32px #18122b55, 0 2px 8px #39375b33;
  transition: background 0.5s, color 0.5s, border 0.5s;
}
.list-view-container .v-card:hover {
  box-shadow: 0 12px 40px #43c6ac55, 0 4px 16px #ffd54f55;
  transform: translateY(-6px) scale(1.04);
}
.list-view-container .v-card .v-card-title {
  font-size: 1.2rem;
  font-weight: 900;
  color: #43c6ac;
  text-shadow: 0 4px 24px #43c6ac33, 0 1px 0 #fffbe7;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
.list-view-container .v-card .v-card-text {
  font-size: 0.95rem;
  color: #fffbe7;
  font-weight: 600;
  text-shadow: 0 1px 8px #00000022;
}
.list-view-container .v-card .v-card-actions {
  justify-content: flex-start;
  gap: 10px;
  padding-top: 12px;
}
.list-view-container .v-card .v-card-actions .v-btn {
  font-weight: bold;
  box-shadow: 0 2px 8px #43c6ac33;
  transition: background 0.2s, color 0.2s;
}
.list-view-container .v-card .v-card-actions .v-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  color: #222 !important;
}
.floating-action-btn {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 10;
  box-shadow: 0 4px 16px #43c6ac55;
  transition: background 0.3s, transform 0.3s;
}
.floating-action-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  transform: translateY(-5px);
}
.quick-view-card {
  background: v-bind('isDark ? "linear-gradient(135deg, #231942 60%, #39375B 100%)" : "linear-gradient(135deg, #fffbe7 60%, #e1f5fe 100%)"');
  border: 2px solid v-bind('isDark ? "#635985" : "#ffd54f"');
  color: v-bind('isDark ? "#fff" : "#333"');
  box-shadow: 0 8px 32px #18122b55, 0 2px 8px #39375b33;
  transition: background 0.5s, color 0.5s, border 0.5s;
}
.quick-view-card .v-card-title {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  font-size: 1.5rem;
  font-weight: 900;
  letter-spacing: 1.5px;
  color: #fff;
  text-shadow: 0 2px 12px #00000033, 0 1px 0 #fffbe7;
  margin-bottom: 12px;
}
.quick-view-card .v-card-text {
  font-size: 1rem;
  color: #fffbe7;
  font-weight: 600;
  text-shadow: 0 1px 8px #00000022;
  margin-bottom: 20px;
}
.quick-view-card .v-card-actions {
  justify-content: flex-start;
  gap: 15px;
}
.quick-view-card .v-card-actions .v-btn {
  font-weight: bold;
  box-shadow: 0 2px 8px #43c6ac33;
  transition: background 0.2s, color 0.2s;
}
.quick-view-card .v-card-actions .v-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  color: #222 !important;
}
.card-image-container {
  position: relative;
  height: 200px;
  overflow: hidden;
}
.card-image {
  transition: transform 0.5s ease-in-out;
}
.card-image-container:hover .card-image {
  transform: scale(1.1);
}
.card-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  opacity: 0;
  transition: opacity 0.3s ease-in-out;
  z-index: 1;
}
.card-image-container:hover .card-overlay {
  opacity: 1;
}
.quick-view-btn, .favorite-btn {
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(5px);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: background 0.3s, border-color 0.3s;
}
.quick-view-btn:hover, .favorite-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
}
.favorite-btn {
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
}
.favorite-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.4);
}
.animated-card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}
.animated-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 12px 40px #43c6ac55, 0 4px 16px #ffd54f55;
}
/* Modern Category Filter Styles */
.category-filter-container {
  position: relative;
  z-index: 10;
}

.category-filter-wrapper {
  position: relative;
}

.modern-category-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  color: #fff !important;
  font-weight: 600;
  letter-spacing: 0.5px;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.4);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  position: relative;
  overflow: hidden;
}

.modern-category-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s;
}

.modern-category-btn:hover::before {
  left: 100%;
}

.modern-category-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 12px 40px rgba(102, 126, 234, 0.6);
  border-color: rgba(255, 255, 255, 0.3);
}

.modern-category-dropdown {
  background: v-bind('isDark ? "linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%)" : "linear-gradient(135deg, #ffffff 0%, #f8f9fa 50%, #e9ecef 100%)"');
  border: 2px solid v-bind('isDark ? "rgba(255, 255, 255, 0.1)" : "rgba(102, 126, 234, 0.2)"');
  color: v-bind('isDark ? "#fff" : "#333"');
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(20px);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  max-height: 600px;
  overflow: hidden;
}

.modern-dropdown-title {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff !important;
  font-weight: 700;
  letter-spacing: 0.5px;
  padding: 20px 24px;
  border-radius: 12px 12px 0 0;
  position: relative;
  overflow: hidden;
}

.modern-dropdown-title::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, transparent 30%, rgba(255, 255, 255, 0.1) 50%, transparent 70%);
  animation: shimmer 2s infinite;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.close-btn {
  background: rgba(255, 255, 255, 0.1) !important;
  border-radius: 50%;
  transition: all 0.3s ease;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  transform: rotate(90deg);
}

.modern-dropdown-content {
  padding: 0;
  max-height: 400px;
  overflow-y: auto;
  scrollbar-width: thin;
  scrollbar-color: v-bind('isDark ? "#667eea #1a1a2e" : "#667eea #f8f9fa"');
}

.modern-dropdown-content::-webkit-scrollbar {
  width: 6px;
}

.modern-dropdown-content::-webkit-scrollbar-track {
  background: v-bind('isDark ? "#1a1a2e" : "#f8f9fa"');
  border-radius: 3px;
}

.modern-dropdown-content::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 3px;
}

.modern-category-list {
  background: transparent;
  padding: 8px;
}

.modern-category-item {
  margin: 4px 0;
  border-radius: 16px;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.modern-category-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.1), transparent);
  transition: left 0.6s;
}

.modern-category-item:hover::before {
  left: 100%;
}

.modern-category-item:hover {
  background: v-bind('isDark ? "rgba(255, 255, 255, 0.08)" : "rgba(102, 126, 234, 0.08)"');
  transform: translateX(8px) scale(1.02);
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.2);
}

.modern-category-item.selected-category-item {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  box-shadow: 0 6px 24px rgba(102, 126, 234, 0.4);
  transform: translateX(4px);
}

.modern-category-avatar {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.modern-category-item:hover .modern-category-avatar {
  transform: scale(1.15) rotate(5deg);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.3);
}

.selected-category-item .modern-category-avatar {
  border-color: rgba(255, 255, 255, 0.4);
  box-shadow: 0 6px 24px rgba(255, 255, 255, 0.3);
}

.modern-category-title {
  font-weight: 600;
  letter-spacing: 0.3px;
  line-height: 1.4;
}

.modern-category-count {
  font-weight: 700;
  font-size: 0.75rem;
  min-width: 28px;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
  transition: all 0.3s ease;
}

.modern-category-item:hover .modern-category-count {
  transform: scale(1.1);
}

.selected-icon {
  animation: checkmark 0.3s ease-in-out;
}

@keyframes checkmark {
  0% { transform: scale(0) rotate(-180deg); opacity: 0; }
  50% { transform: scale(1.2) rotate(-90deg); opacity: 0.7; }
  100% { transform: scale(1) rotate(0deg); opacity: 1; }
}

.modern-dropdown-actions {
  background: v-bind('isDark ? "rgba(255, 255, 255, 0.05)" : "rgba(102, 126, 234, 0.05)"');
  border-top: 1px solid v-bind('isDark ? "rgba(255, 255, 255, 0.1)" : "rgba(102, 126, 234, 0.1)"');
  padding: 16px 24px;
}

.show-all-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  color: #fff !important;
  font-weight: 600;
  letter-spacing: 0.5px;
  border: 2px solid rgba(255, 255, 255, 0.2);
  transition: all 0.3s ease;
}

.show-all-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
  border-color: rgba(255, 255, 255, 0.4);
}

.done-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  color: #fff !important;
  font-weight: 600;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
}

.done-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
}

/* Responsive Design for Modern Category Filter */
@media (max-width: 768px) {
  .modern-category-btn {
    min-width: 160px;
    height: 48px;
  }
  
  .modern-category-dropdown {
    max-width: 320px;
    min-width: 280px;
  }
  
  .modern-dropdown-title {
    padding: 16px 20px;
  }
  
  .modern-dropdown-actions {
    padding: 12px 20px;
  }
}

/* Mega Menu Styles */
.mega-category-bar {
  position: relative;
  z-index: 10;
}

.mega-category-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  color: #fff !important;
  font-weight: 600;
  letter-spacing: 0.5px;
  box-shadow: 0 8px 32px rgba(102, 126, 234, 0.4);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  position: relative;
  overflow: hidden;
}

.mega-category-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.6s;
}

.mega-category-btn:hover::before {
  left: 100%;
}

.mega-category-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 0 12px 40px rgba(102, 126, 234, 0.6);
  border-color: rgba(255, 255, 255, 0.3);
}

.mega-menu-dropdown {
  background: v-bind('isDark ? "linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%)" : "linear-gradient(135deg, #ffffff 0%, #f8f9fa 50%, #e9ecef 100%)"');
  border: 2px solid v-bind('isDark ? "rgba(255, 255, 255, 0.1)" : "rgba(102, 126, 234, 0.2)"');
  color: v-bind('isDark ? "#fff" : "#333"');
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(20px);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  max-height: 600px;
  overflow: hidden;
}

.mega-main-list {
  padding: 20px;
  border-right: 1px solid v-bind('isDark ? "rgba(255, 255, 255, 0.1)" : "rgba(102, 126, 234, 0.2)"');
}

.mega-main-item {
  margin-bottom: 15px;
  padding: 10px 15px;
  border-radius: 12px;
  transition: all 0.3s ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 15px;
}

.mega-main-item:hover {
  background: v-bind('isDark ? "rgba(255, 255, 255, 0.08)" : "rgba(102, 126, 234, 0.08)"');
  transform: translateX(5px);
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.2);
}

.mega-main-item.active-main {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  box-shadow: 0 6px 24px rgba(102, 126, 234, 0.4);
  transform: translateX(4px);
}

.mega-main-avatar {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.mega-main-item:hover .mega-main-avatar {
  transform: scale(1.15) rotate(5deg);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.3);
}

.mega-sub-list {
  padding: 20px;
}

.mega-sub-item {
  margin-bottom: 10px;
  padding: 10px 15px;
  border-radius: 12px;
  transition: all 0.3s ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 15px;
}

.mega-sub-item:hover {
  background: v-bind('isDark ? "rgba(255, 255, 255, 0.08)" : "rgba(102, 126, 234, 0.08)"');
  transform: translateX(5px);
  box-shadow: 0 4px 20px rgba(102, 126, 234, 0.2);
}

.mega-sub-avatar {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.mega-sub-item:hover .mega-sub-avatar {
  transform: scale(1.15) rotate(5deg);
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.3);
}

.mega-sub-title {
  font-weight: 600;
  letter-spacing: 0.3px;
  line-height: 1.4;
}

/* Responsive Mega Menu */
@media (max-width: 1024px) {
  .mega-menu-dropdown {
    max-width: 90vw;
    min-width: 90vw;
    max-height: 80vh;
    overflow-y: auto;
  }

  .mega-main-list, .mega-sub-list {
    padding: 15px;
  }

  .mega-main-item, .mega-sub-item {
    padding: 8px 12px;
    gap: 12px;
  }

  .mega-main-title, .mega-sub-title {
    font-size: 1rem;
  }

  .mega-sub-avatar {
    width: 28px;
    height: 28px;
  }
}

@media (max-width: 768px) {
  .mega-menu-dropdown {
    max-width: 95vw;
    min-width: 95vw;
    max-height: 70vh;
  }

  .mega-main-list, .mega-sub-list {
    padding: 10px;
  }

  .mega-main-item, .mega-sub-item {
    padding: 6px 10px;
    gap: 10px;
  }

  .mega-main-title, .mega-sub-title {
    font-size: 0.9rem;
  }

  .mega-sub-avatar {
    width: 24px;
    height: 24px;
  }
}
</style> 