<template>
  <v-app>
    <v-main>
      <v-container class="d-flex align-center justify-center" style="min-height: 100vh;">
        <v-card class="pa-8" max-width="400" elevation="10" rounded="xl">
          <div class="text-center mb-6">
            <v-avatar size="64" class="mb-2" color="primary">
              <v-icon size="36" color="white">mdi-account-plus</v-icon>
            </v-avatar>
            <h2 class="text-h5 font-weight-bold mb-1">Sign Up</h2>
            <div class="text-body-2 mb-2">Create a new account</div>
          </div>
          <v-form @submit.prevent="handleRegister">
            <v-text-field
              v-model="name"
              label="Name Surname"
              prepend-inner-icon="mdi-account"
              required
              class="mb-4"
              autocomplete="name"
            />
            <v-text-field
              v-model="email"
              label="E-mail"
              type="email"
              prepend-inner-icon="mdi-email"
              required
              class="mb-4"
              autocomplete="username"
            />
            <v-text-field
              v-model="password"
              label="Password"
              :type="showPassword ? 'text' : 'password'"
              prepend-inner-icon="mdi-lock"
              :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              @click:append-inner="showPassword = !showPassword"
              required
              class="mb-4"
              autocomplete="new-password"
            />
            <v-text-field
              v-model="confirmPassword"
              label="Repeat Password"
              :type="showConfirmPassword ? 'text' : 'password'"
              prepend-inner-icon="mdi-lock-check"
              :append-inner-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
              @click:append-inner="showConfirmPassword = !showConfirmPassword"
              required
              class="mb-6"
              autocomplete="new-password"
            />
            <v-btn 
              color="primary" 
              type="submit" 
              block 
              size="large" 
              class="mb-2" 
              rounded="xl"
              :loading="loading"
            >
              Sign Up
            </v-btn>
            <v-btn variant="text" block @click="$router.push('/login')" class="mb-4">
              Already have an account? Log In
            </v-btn>
          </v-form>
          <v-alert
            v-if="error"
            type="error"
            variant="tonal"
            class="mb-4"
          >
            {{ error }}
          </v-alert>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { register } from './api/productApi';

const router = useRouter();
const name = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const loading = ref(false);
const error = ref('');
const showPassword = ref(false);
const showConfirmPassword = ref(false);

async function handleRegister() {
  if (password.value !== confirmPassword.value) {
    error.value = 'Şifreler eşleşmiyor';
    return;
  }

  try {
    loading.value = true;
    error.value = '';
    await register({
      name: name.value,
      email: email.value,
      password: password.value
    });
    router.push('/login');
  } catch (err) {
    error.value = err.response?.data?.message || 'Kayıt başarısız';
  } finally {
    loading.value = false;
  }
}
</script>

<style>
.v-card {
  border-radius: 24px !important;
}

.v-btn {
  text-transform: none !important;
  font-weight: 600 !important;
  letter-spacing: 0.5px !important;
}

.v-btn--rounded {
  border-radius: 16px !important;
}

.v-avatar {
  border: 4px solid var(--v-primary-lighten2);
}

.v-text-field :deep(.v-field__outline__start) {
  border-radius: 12px 0 0 12px !important;
}

.v-text-field :deep(.v-field__outline__end) {
  border-radius: 0 12px 12px 0 !important;
}
</style> 