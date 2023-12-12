import {createRouter , createWebHistory} from "vue-router" 
import UserLogin from "../views/UserLogin.vue"
import UserSignUp from "../views/UserSignUp.vue"
const router  = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes:[
        {
            path: "/",
            name: "userlogin",
            component: UserLogin
        },
        {
            path: "/usersignup",
            name: "usersignup",
            component: UserSignUp
        }
    ]
})

export default router