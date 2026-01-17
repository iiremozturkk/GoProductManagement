import { createRouter, createWebHistory } from 'vue-router';
import Login from './Login.vue';
import Register from './Register.vue';
import ProductGallery from './ProductGallery.vue';
import ProductDetail from './ProductDetail.vue';
import Cart from './Cart.vue';
import Profile from './Profile.vue';
import ProductForm from './components/ProductForm.vue';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
  },
  {
    path: '/gallery',
    name: 'ProductGallery',
    component: ProductGallery,
    meta: { requiresAuth: true },
  },
  {
    path: '/product/new',
    name: 'AddProduct',
    component: ProductForm,
    meta: { requiresAuth: true },
  },
  {
    path: '/product/:id',
    name: 'ProductDetail',
    component: ProductDetail,
    props: true,
    meta: { requiresAuth: true },
  },
  {
    path: '/product/:id/edit',
    name: 'EditProduct',
    component: ProductForm,
    props: route => ({
      isEditing: true
    }),
    meta: { requiresAuth: true },
  },
  {
    path: '/cart',
    name: 'Cart',
    component: Cart,
    meta: { requiresAuth: true },
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    meta: { requiresAuth: true },
  },
  {
    path: '/',
    redirect: '/gallery',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
  if (to.meta.requiresAuth && !isLoggedIn) {
    next('/login');
  } else if ((to.path === '/login' || to.path === '/register') && isLoggedIn) {
    next('/product/new');
  } else {
    next();
  }
});

export default router;