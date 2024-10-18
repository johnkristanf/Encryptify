<script lang="ts" setup>
import { ref, onMounted } from 'vue';

// Define an interface for User
interface User {
  email: string;
  password: string;
  salt: string;
  pepper: string;
  created_at: string; // Use string for dates fetched from API
}

// Define a reactive variable to store fetched users data
const users = ref<User[]>([]); // The users array will now have type User[]

// Function to fetch data from the API
const fetchUsers = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/users');
    if (!response.ok) {
      throw new Error('Failed to fetch users');
    }
    // Parse the JSON response and assign it to the users array
    const data: User[] = await response.json(); // Explicitly declare that the response will be of type User[]
    users.value = data; // Assign the fetched data to the users array
  } catch (error) {
    console.error(error);
  }
};

// Fetch data when the component is mounted
onMounted(() => {
  fetchUsers();
});
</script>

<template>
  <div class="w-full h-screen flex justify-center pb-5 ">
    <div class="w-[90%] mt-4">
      <h1 class="font-semibold my-5 text-4xl">IT412 MIDTERM PROJECT</h1>
      <div class="w-full relative overflow-x-auto">
        <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
          <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="px-6 py-3">Email</th>
              <th scope="col" class="px-6 py-3">Password</th>
              <th scope="col" class="px-6 py-3">Salt</th>
              <th scope="col" class="px-6 py-3">Pepper</th>
              <th scope="col" class="px-6 py-3">Date Registered</th>
            </tr>
          </thead>
          <tbody>
            <!-- Loop through users array and display each user in the table -->
            <tr
              v-for="(user, index) in users"
              :key="index"
              class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
            >
              <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                {{ user.email }}
              </td>
              <td class="px-6 py-4">{{ user.password }}</td>
              <td class="px-6 py-4">{{ user.salt }}</td>
              <td class="px-6 py-4">{{ user.pepper }}</td>
              <td class="px-6 py-4">{{ new Date(user.created_at).toLocaleString() }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
