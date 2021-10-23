<template>
    <ion-page>
        <ion-tabs>
            <ion-router-outlet></ion-router-outlet>
            <ion-tab-bar slot="bottom">
            
                <ion-tab-button tab="login" href="/auth/login">
                    <ion-icon :icon="logInSharp" />
                    <ion-label>Login</ion-label>
                </ion-tab-button>
              
                <ion-tab-button tab="createAccount" href="/auth/createAccount">
                    <ion-icon :icon="personAddSharp" />
                    <ion-label>Sign up</ion-label>
                </ion-tab-button>
                
            </ion-tab-bar>
        </ion-tabs>
    </ion-page>
</template>

<script>
import { IonTabBar, IonTabButton, IonTabs, IonLabel, IonIcon, IonPage, IonRouterOutlet } from '@ionic/vue';
import { logInSharp, personAddSharp } from 'ionicons/icons';
import { mapGetters } from "vuex";

export default {
    name: 'AuthTabs',
    components: { IonLabel, IonTabs, IonTabBar, IonTabButton, IonIcon, IonPage, IonRouterOutlet },
    computed: {
        ...mapGetters(["isAuthenticated"])
    },
    watch: {
        isAuthenticated(val) {
            // automatically redirect user when logged in
            if (val)
                this.$router.push("/tabs/");
        }
    },
    mounted() {

        // try to login from token in localStorage
        this.$store.dispatch("restoreLogin");

        // restore theme from localStorage (dark by default)
        const colorTheme = localStorage.getItem("theme") || "dark";
        document.body.setAttribute("color-theme", colorTheme);

    },
    setup() {
        return { logInSharp, personAddSharp };
    }
}
</script>
