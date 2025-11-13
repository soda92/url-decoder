<template>
  <div id="app">
    <h1>WWW Form URL Encoded Decoder</h1>
    <textarea v-model="encodedData" placeholder="Enter www-form-urlencoded data here"></textarea>
    <input type="text" v-model="searchTerm" placeholder="Search by field name" />
    <div class="button-container">
      <button @click="showJson = !showJson">{{ showJson ? 'Hide JSON' : 'View as JSON' }}</button>
      <button v-if="showJson" @click="copyJson">Copy JSON</button>
    </div>
    <pre v-if="showJson">{{ jsonOutput }}</pre>
    <table>
      <thead>
        <tr>
          <th>Field Name</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in filteredDecodedData" :key="index">
          <td>{{ item.key }}</td>
          <td>{{ item.value }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const encodedData = ref('');
const searchTerm = ref('');
const showJson = ref(false);

const decodedData = computed(() => {
  if (!encodedData.value) {
    return [];
  }
  try {
    const searchParams = new URLSearchParams(encodedData.value);
    const data = [];
    for (const [key, value] of searchParams.entries()) {
      data.push({ key, value });
    }
    return data;
  } catch (error) {
    console.error('Error decoding data:', error);
    return [];
  }
});

const jsonOutput = computed(() => {
  const obj: { [key: string]: string | string[] } = {};
  for (const { key, value } of decodedData.value) {
    if (obj.hasOwnProperty(key)) {
      if (Array.isArray(obj[key])) {
        (obj[key] as string[]).push(value);
      } else {
        obj[key] = [obj[key] as string, value];
      }
    } else {
      obj[key] = value;
    }
  }
  return JSON.stringify(obj, null, 2);
});

const copyJson = () => {
  navigator.clipboard.writeText(jsonOutput.value)
    .then(() => {
      alert('JSON copied to clipboard!');
    })
    .catch(err => {
      console.error('Failed to copy JSON: ', err);
      alert('Failed to copy JSON. See console for details.');
    });
};

const filteredDecodedData = computed(() => {
  if (!searchTerm.value) {
    return decodedData.value;
  }
  return decodedData.value.filter(item =>
    item.key.toLowerCase().includes(searchTerm.value.toLowerCase())
  );
});
</script>

<style scoped>
#app {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
  font-family: sans-serif;
}

h1 {
  text-align: center;
  margin-bottom: 2rem;
}

textarea {
  width: 100%;
  height: 150px;
  margin-bottom: 1rem;
  padding: 0.5rem;
  font-size: 1rem;
}

input {
  width: 100%;
  padding: 0.5rem;
  font-size: 1rem;
  margin-bottom: 1rem;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 0.5rem;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}

.button-container {
  margin-bottom: 1rem;
  display: flex;
  gap: 1rem;
}

button {
  padding: 0.5rem 1rem;
  font-size: 1rem;
  cursor: pointer;
}

pre {
  background-color: #f2f2f2;
  padding: 1rem;
  border: 1px solid #ddd;
  white-space: pre-wrap;
  word-wrap: break-word;
  margin-bottom: 1rem;
}
</style>