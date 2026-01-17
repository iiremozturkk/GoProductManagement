<template>
  <v-card elevation="2">
    <v-card-title>Products</v-card-title>
    <v-data-table
      :headers="headers"
      :items="products"
      class="elevation-1"
      item-key="id"
      :items-per-page="5"
      no-data-text="No products found."
    >
      <template #item.actions="{ item }">
        <v-btn icon="mdi-pencil" color="primary" @click="$emit('edit', item)"></v-btn>
        <v-btn icon="mdi-delete" color="error" @click="deleteProduct(item.id)"></v-btn>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup>
import { deleteProduct as apiDeleteProduct } from '../api/productApi';

const props = defineProps({
  products: Array
});
const emit = defineEmits(['edit', 'deleted']);

const headers = [
  { title: 'ID', key: 'id', align: 'start', sortable: false },
  { title: 'Name', key: 'name' },
  { title: 'Description', key: 'description' },
  { title: 'Price', key: 'price' },
  { title: 'Stock', key: 'stock' },
  { title: 'Actions', key: 'actions', sortable: false }
];

async function deleteProduct(id) {
  if (confirm('Bu ürünü silmek istediğinize emin misiniz?')) {
    await apiDeleteProduct(id);
    emit('deleted');
  }
}
</script> 