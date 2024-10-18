
<script setup lang="ts">
    import LoginFrom from '../components/LoginForm.vue';
</script>

<template>
    <div class="h-screen w-full flex justify-around items-center font-semibold bg-blue-300">

        <div class="flex flex-col gap-5">

            <div class="flex items-center">
                <img src="/img/encryptify_logo.png" alt="Encyptify Logo" width="120" height="120">
                <h1 class="text-white text-7xl">Encryptify</h1>
            </div>

            <h1 class="text-md opacity-75">
                "Welcome to Encryptify, the ultimate solution
                for safeguarding your sensitive information. <br>

                Our cutting-edge platform uses advanced hashing 
                techniques combined with salt and pepper functionalities
            </h1>
        </div>

        <LoginFrom />

        
    </div>
</template>


<style>
    .background-image {
        background-image: url('/img/encryptify_bg.jpg');
        background-size: cover;
        background-position: center;
        background-repeat: no-repeat;        
    }
</style>
