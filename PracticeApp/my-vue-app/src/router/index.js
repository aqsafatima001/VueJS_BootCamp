import {createRouter , createWebHistory} from "vue-router" 
import UserLogin from "../views/UserLogin.vue"
const router  = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes:[
        {
            path: "/",
            name: "userlogin",
            component: UserLogin
        }
        // {
        //     path: "/quiz/:id",
        //     name: "quiz",
        //     component: QuizView
        // }
    ]
})

export default router