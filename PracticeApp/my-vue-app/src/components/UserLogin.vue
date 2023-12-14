<script>
import  axios from 'axios';

export default {
    data() {
        return {
            formData: {
                username: '',
                password: '',
            },
        };
    },
    methods: {
        submitLoginForm() {
            axios.post('/api/login', this.formData)
                .then(response => {
                    // Check the response from the backend
                    if (response.data.success) {
                        // Redirect to /clusterhome on successful authentication
                        this.$router.push('/clusterhome');
                    } else {
                        // Handle authentication failure
                        alert('Authentication failed. Please check your credentials.');
                    }
                })
                .catch(error => {
                    // Handle errors
                    console.error('Error during login', error);
                });
        },
    },
};
</script>


<template>
    <div class="container">

        <h1>User Login</h1>
        <!-- <form action="/api/login" method="post"> -->
        <form @submit.prevent="submitLoginForm" action="/api/login" method="post" >
            <div>
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
            </div>
            <div>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
            </div>
            <div class="button-container">
                <button @click="redirectToGolangBackend" class="Login-button" type="submit">Log in</button>
                <!-- <router-link to="/clusterhome" class="Login-button">Log in</router-link> -->
            </div>
            <hr>
        </form>

        <div class="button-container">
            <router-link to="/usersignup" class="SignUp-button">Create new Account</router-link>
        </div>
    </div>
</template>

<style scoped>
* {
    margin: 0;
    padding: 0;
}

.container {
    margin: 0 auto;
    max-width: 400px;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
    position: absolute;
    /* or use 'fixed' if you want it to stay in the center even when scrolling */
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 1000px;
}

h1 {
    font-size: 48px;
    text-align: center;
    font-family: sans-serif;
    font-weight: bold;
    color: #27496D;
}

label {
    display: block;
    margin-top: 5px;
    font-family: sans-serif;
}

input[type="text"],
input[type="password"] {
    padding: 5px;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;

}

/* Default button style */
.button-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    /* Horizontally center content */
    align-items: center;
    /* Vertically center content */
}

.Login-button {
    padding: 10px 20px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    background-color: #27496D;
    /* Change to your preferred background color */
    color: #fff;
    /* Change to your preferred text color */
    cursor: pointer;
    transition: background-color 0.3s ease, color 0.3s ease;
    margin-top: 25px;
    width: 100%
}

/* Hover effect */
.Login-button:hover {
    background-color: #0056b3;
    /* Change to your preferred hover background color */
}

.SignUp-button {
    padding: 10px 20px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    background-color: #7abdf4;
    /* Change to your preferred background color */
    color: #fff;
    /* Change to your preferred text color */
    cursor: pointer;
    transition: background-color 0.3s ease, color 0.3s ease;
    margin-top: 25px;
    width: 100%;
    text-decoration: none;
    text-align: center;
}

hr {
    margin-top: 25px;
    opacity: 0.7;
}



label {
    display: block;
}
</style>