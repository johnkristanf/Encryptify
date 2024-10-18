<script lang="ts" setup>
import { ref } from 'vue';

// Reactive variables for the form inputs
const email = ref('');
const password = ref('');
const errorMessage = ref(''); // Store any error message from the request
const successMessage = ref(''); // Store success message if login is successful

// Handle form submission
const handleSubmit = async () => {
    errorMessage.value = '';
    successMessage.value = '';

    const payload = {
      email: email.value,
      password: password.value
    };

    console.log("login payload: ", payload)

    try {
        // Send POST request using fetch API
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload),
        });

        if (response.ok) {
          const data = await response.json();
          successMessage.value = 'Login successful!';
          console.log('Login successful:', data);
          window.location.href = '/dashboard';

        } else {
          errorMessage.value = 'The credentials you entered are incorrect. Please check your details and try again.';
        }
    } catch (error) {
        console.error("Error in Login: ", error)
        errorMessage.value = 'An error occurred during login. Please try again.';
    }
};
</script>

<template>
  <section class="bg-gray-50 dark:bg-gray-900 h-[80%] w-[30%] flex items-center">
    <div class="w-full flex flex-col items-center justify-center px-6 py-8">
      <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
            Sign in to your account
          </h1>

          <!-- Display success or error message here, below the heading -->
          <p v-if="errorMessage" class="text-red-500 text-sm">{{ errorMessage }}</p>
          <p v-if="successMessage" class="text-green-500 text-sm">{{ successMessage }}</p>

          <!-- Form starts here -->
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleSubmit">
            <div>
              <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
              <input
                type="email"
                name="email"
                id="email"
                v-model="email"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="name@company.com"
                required
              />
            </div>

            <div>
              <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
              <input
                type="password"
                name="password"
                id="password"
                v-model="password"
                placeholder="••••••••"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
            </div>

            <button
              type="submit"
              class="w-full text-white bg-blue-600 hover:opacity-75 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Sign in
            </button>

            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Don’t have an account yet?
              <router-link to="/register" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</router-link>
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
