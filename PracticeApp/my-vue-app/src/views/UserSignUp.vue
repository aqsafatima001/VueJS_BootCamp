<script setup>
import Header from "../components/header.vue"
import Footer from "../components/Footer.vue"
</script>

<template>
    <Header />

    <div class="container">
        <h1>User Sign Up</h1>
        <form action="signup.php" method="post"> <!-- Replace 'signup.php' with your server-side script -->
            <!-- User Information -->
            <div class="form-group">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username" required>
            </div>
            <!-- User Information -->

            <div class="name-group">
                <div class="form-group">
                    <label for="title">Title:</label>
                    <select id="title" name="title">
                        <option value="Mr.">Mr.</option>
                        <option value="Ms.">Ms.</option>
                        <option value="Mrs.">Mrs.</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="firstName">First Name:</label>
                    <input type="text" id="firstName" name="firstName" required>
                </div>
                <div class="form-group">
                    <label for="middleName">Middle Name:</label>
                    <input type="text" id="middleName" name="middleName">
                </div>
                <div class="form-group">
                    <label for="lastName">Last Name:</label>
                    <input type="text" id="lastName" name="lastName" required>
                </div>
            </div>

            <!-- Password Information -->
            <div class="form-group">
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required>
            </div>

            <!-- Combine Password and Confirm Password in one line -->
            <div class="password-group">
                <div class="form-group">
                    <label for="confirmPassword">Confirm Password:</label>
                    <input type="password" id="confirmPassword" name="confirmPassword" required>
                </div>
            </div>

            <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" id="email" name="email" required>
            </div>
            <!-- Add more user fields here -->

            <!-- Thesis Information -->
            <div class="form-group">
                <label for="thesisTitle">Thesis Title:</label>
                <input type="text" id="thesisTitle" name="thesisTitle" required>
            </div>
            <!-- Add more thesis fields here -->

            <!-- Services Information -->
            <div class="form-group">
                <label>Services:</label><br>
                <input type="checkbox" id="hadoop" name="hadoop"> Hadoop<br>
                <input type="checkbox" id="hive" name="hive"> Hive<br>
                <input type="checkbox" id="Pig" name="Pig"> Pig<br>
                <input type="checkbox" id="Spark" name="Spark"> Spark<br>
                <input type="checkbox" id="Sqoop" name="Sqoop"> Sqoop<br>
                <input type="checkbox" id="MySQL" name="MySQL"> MySQL<br>
                <input type="checkbox" id="Zookeeper" name="Zookeeper"> Zookeeper<br>
                <input type="checkbox" id="Hbase" name="Hbase"> Hbase<br>
                <input type="checkbox" id="Kafka" name="Kafka"> KAfka<br>
                <input type="checkbox" id="Flume" name="Flume"> Flume<br>
                <input type="checkbox" id="Storm" name="Storm"> Storm<br>
                <!-- Add more service checkboxes here -->
            </div>

            <!-- Affiliation Information -->
            <div class="form-group">
                <label for="xyzid">XYZ ID:</label>
                <input type="text" id="xyzid" name="xyzid">
            </div>
            <!-- Add more affiliation fields here -->

            <!-- Account Duration Information -->
            <div class="form-group">
                <label for="startDate">Start Date:</label>
                <input type="date" id="startDate" name="startDate">
            </div>
            <div class="form-group">
                <label for="endDate">End Date:</label>
                <input type="date" id="endDate" name="endDate">
            </div>

            <button type="submit">Sign Up</button>
        </form>
    </div>

    <Footer />
</template>


<style scoped>
/* Styles for the container */

* {
    margin: 0;
    padding: 0;
}

.container {

    margin: 0 auto;
    max-width: 800px;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
    margin-top: 20px;
    margin-bottom: 100px;
    position: relative;
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
    font-family: sans-serif;

}

/* Add styles for the password-group */
.password-group {
    display: flex;
    justify-content: space-between;
}

/* Adjust the width of the password fields */
.password-group .form-group {
    flex: 1;
    margin-right: 10px; /* Add margin between password and confirm password */
}

/* Set equal width for password and confirm password fields */
.password-group .form-group:last-child {
    margin-right: 0;
}

/* Styles for the form elements */
.form-group {
    margin-bottom: 15px;
}

.name-group {
    display: flex;
    justify-content: space-between;
}
input[type="text"],
input[type="email"],
select,
input[type="date"],
input[type="password"] {
    padding: 5px;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;
    border: 2px solid rgb(207, 206, 206);

}

/* Style the checkbox label */
.checkbox-label {
    display: inline-block;
    margin-right: 10px;
}

/* Style the Submit button */
button[type="submit"] {
    background-color: #007BFF;
    color: #fff;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button[type="submit"]:hover {
    background-color: #0056b3;
}

</style>