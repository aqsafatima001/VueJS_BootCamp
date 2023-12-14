import {createRouter , createWebHistory} from "vue-router" 
import UserLogin from "../views/UserLogin.vue"
import UserSignUp from "../views/UserSignUp.vue"
import ClusterHome from "../views/ClusterHome.vue"
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
        },
        {
            path: "/clusterhome",
            name: "Clusterhome",
            component: ClusterHome
        },
        {
            path: "/api/login",
            name: "Clusterhome",
            component: ClusterHome
        }
        

    ]
})

export default router