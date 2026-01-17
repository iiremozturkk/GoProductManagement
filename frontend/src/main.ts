import { createApp } from 'vue';
import App from './App.vue';
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import '@mdi/font/css/materialdesignicons.css';
import router from './router';

const vuetify = createVuetify({
  theme: {
    defaultTheme: 'myCustomTheme',
    themes: {
      myCustomTheme: {
        dark: false,
        colors: {
          primary: '#1976D2',
          secondary: '#FF9800',
          accent: '#E040FB',
          error: '#FF5252',
          info: '#2196F3',
          success: '#43A047',
          warning: '#FFC107',
          background: '#f0f4ff',
          surface: '#ffffff',
        },
      },
      myDarkTheme: {
        dark: true,
        colors: {
          primary: '#90caf9',
          secondary: '#ffb74d',
          accent: '#ce93d8',
          error: '#ef9a9a',
          info: '#64b5f6',
          success: '#81c784',
          warning: '#ffd54f',
          background: '#181a20',
          surface: '#232634',
        },
      },
    },
  },
  icons: {
    defaultSet: 'mdi',
  },
});

createApp(App)
  .use(vuetify)
  .use(router)
  .mount('#app'); 