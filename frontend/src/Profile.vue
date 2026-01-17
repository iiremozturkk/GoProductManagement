<template>
  <v-app>
    <v-main class="profile-bg">
      <v-container class="py-10 d-flex align-center justify-center position-relative" style="min-height: 100vh;">
        <div class="profile-bg-shapes">
          <div class="shape1"></div>
          <div class="shape2"></div>
          <div class="shape3"></div>
          <div class="shape4"></div>
          <div class="shape5"></div>
          <div class="shape6"></div>
          <div class="shape7"></div>
        </div>
        <v-card class="profile-card pa-14" max-width="800" elevation="12" rounded="2xl">
          <v-card-title class="text-center mb-4 d-flex flex-column align-center">
            <v-avatar size="80" class="mb-2" color="primary">
              <v-icon size="56" color="white">mdi-account-circle</v-icon>
            </v-avatar>
            <div class="profile-title mb-1">User Profile</div>
            <div class="profile-subtitle">Update your information</div>
          </v-card-title>
          <v-card-text>
            <v-form @submit.prevent="saveProfile">
              <v-text-field v-model="form.name" label="Name" prepend-inner-icon="mdi-account" required rounded="xl" variant="solo"></v-text-field>
              <v-text-field v-model="form.email" label="Email" prepend-inner-icon="mdi-email" required type="email" rounded="xl" variant="solo" class="profile-input-wide"></v-text-field>
              <v-text-field v-model="form.password" label="New Password" prepend-inner-icon="mdi-lock" type="password" rounded="xl" variant="solo" class="profile-input-wide"></v-text-field>
              <v-btn color="primary" class="mt-6 profile-save-btn" type="submit" block rounded="xl" size="large">Save Changes</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const user = JSON.parse(localStorage.getItem('user') || '{"name":"John Doe","email":"john@example.com"}');
const form = ref({
  name: user.name,
  email: user.email,
  password: ''
});

function saveProfile() {
  localStorage.setItem('user', JSON.stringify({
    name: form.value.name,
    email: form.value.email
  }));
  alert('Profile updated!');
  router.push('/gallery');
}
</script>

<style>
.profile-bg {
  background: linear-gradient(135deg, #f8ffae 0%, #43c6ac 100%);
  min-height: 100vh;
}
.position-relative {
  position: relative;
}
.profile-bg-shapes {
  position: absolute;
  top: 0; left: 0; width: 100vw; height: 100vh;
  z-index: 0;
  pointer-events: none;
}
.profile-bg-shapes::before {
  content: '';
  position: absolute;
  top: -80px; left: -120px;
  width: 340px; height: 340px;
  background: radial-gradient(circle at 60% 40%, #b39ddb 0%, #392f5a 100%);
  opacity: 0.25;
  border-radius: 50%;
}
.profile-bg-shapes::after {
  content: '';
  position: absolute;
  bottom: -100px; right: -100px;
  width: 260px; height: 260px;
  background: radial-gradient(circle at 40% 60%, #43c6ac 0%, #18122b 100%);
  opacity: 0.18;
  border-radius: 50%;
}
/* Ekstra şekiller */
.profile-bg-shapes .shape1 {
  position: absolute;
  top: 6%; left: 4vw;
  width: 160px; height: 160px;
  background: radial-gradient(circle at 60% 40%, #b39ddb 0%, #392f5a 100%);
  opacity: 0.18;
  border-radius: 50%;
  z-index: 0;
}
.profile-bg-shapes .shape2 {
  position: absolute;
  top: 18%; right: 8vw;
  width: 120px; height: 120px;
  background: radial-gradient(circle at 40% 60%, #43c6ac 0%, #b39ddb 100%);
  opacity: 0.13;
  border-radius: 50%;
  z-index: 0;
}
.profile-bg-shapes .shape3 {
  position: absolute;
  bottom: 10%; left: 10vw;
  width: 110px; height: 110px;
  background: linear-gradient(120deg, #b39ddb 0%, #635985 100%);
  opacity: 0.12;
  border-radius: 50%;
  z-index: 0;
}
.profile-bg-shapes .shape4 {
  position: absolute;
  bottom: 18%; right: 12vw;
  width: 180px; height: 80px;
  background: linear-gradient(120deg, #9575cd 0%, #392f5a 100%);
  opacity: 0.10;
  border-radius: 40% 60% 60% 40% / 60% 40% 60% 40%;
  z-index: 0;
}
.profile-bg-shapes .shape5 {
  position: absolute;
  top: 60%; left: 60vw;
  width: 90px; height: 90px;
  background: radial-gradient(circle at 40% 60%, #43c6ac 0%, #b39ddb 100%);
  opacity: 0.10;
  border-radius: 50%;
  z-index: 0;
}
.profile-bg-shapes .shape6 {
  position: absolute;
  top: 40%; right: 2vw;
  width: 60px; height: 60px;
  background: radial-gradient(circle at 60% 40%, #635985 0%, #392f5a 100%);
  opacity: 0.13;
  border-radius: 50%;
  z-index: 0;
}
.profile-bg-shapes .shape7 {
  position: absolute;
  bottom: 2vw; right: 2vw;
  width: 120px; height: 120px;
  background: radial-gradient(circle at 60% 40%, #b39ddb 0%, #635985 100%);
  opacity: 0.16;
  border-radius: 50%;
  z-index: 0;
}
.profile-card {
  background: #fffbe7;
  border: 2.5px solid #ffd54f;
  box-shadow: 0 16px 64px #43c6ac33, 0 4px 24px #ffd54f33;
  border-radius: 48px;
  transition: background 0.5s, border 0.5s;
  z-index: 1;
}
.profile-title {
  font-size: 1.6rem;
  font-weight: bold;
  color: #43c6ac;
  letter-spacing: 1px;
  margin-bottom: 2px;
}
.profile-subtitle {
  font-size: 1.05rem;
  color: #888;
  margin-bottom: 2px;
}
.profile-save-btn {
  font-weight: bold;
  letter-spacing: 1px;
  box-shadow: 0 2px 8px #43c6ac33;
  transition: background 0.2s, color 0.2s;
}
.profile-save-btn:hover {
  background: linear-gradient(90deg, #43c6ac 0%, #ffd54f 100%);
  color: #222 !important;
}
.profile-input-wide {
  max-width: 520px;
  width: 100%;
  font-size: 1.15rem;
  margin-left: auto;
  margin-right: auto;
}
@media (max-width: 600px) {
  .profile-card {
    padding: 16px !important;
    border-radius: 20px;
    max-width: 98vw;
  }
  .profile-title {
    font-size: 1.1rem;
  }
}
/* Dark mode desteği */
body[data-theme='dark'] .profile-bg {
  background: linear-gradient(135deg, #18122B 0%, #392F5A 60%, #635985 100%) !important;
}
body[data-theme='dark'] .profile-card {
  background: #231942 !important;
  border: 2px solid #635985 !important;
  color: #fff !important;
}
body[data-theme='dark'] .profile-title {
  color: #b39ddb !important;
}
body[data-theme='dark'] .profile-subtitle {
  color: #bdbdbd !important;
}
</style> 