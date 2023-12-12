<template>
    <header>
        <nav>
            <ul>
                <li><a href="#">Home</a></li>
                <li><a href="#">About</a></li>
                <li><a href="#">Contact Us</a></li>
            </ul>
        </nav>
    </header>
</template>

<style scoped>
/* Reset some default styles and apply basic styling to the header */
body {
    margin: 0;
    padding: 0;
    font-family: Arial, sans-serif;
}

header {
    background-color: #27496D; /* Change to your preferred header background color */
    color: #fff; /* Change to your preferred text color */
    padding: 10px 0;
}

/* Style the navigation menu */
nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
    display: flex;
}

nav li {
    margin-right: 20px; /* Adjust spacing between menu items */
}

nav a {
    text-decoration: none;
    color: #fff; /* Change to your preferred link color */
    font-weight: bold;
}

/* Style navigation links on hover */
nav a:hover {
    color: #ff6600; /* Change to your preferred hover color */
}

</style>