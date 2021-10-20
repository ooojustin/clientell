<template>
    <ion-page>
        <ion-tabs>
            <ion-router-outlet></ion-router-outlet>
            <ion-tab-bar slot="bottom">
            
                <ion-tab-button tab="home" href="/tabs/home">
                    <ion-icon :icon="homeSharp" />
                    <ion-label>Home</ion-label>
                </ion-tab-button>
              
                <ion-tab-button tab="tab2" href="/tabs/tab2">
                    <ion-icon :icon="ellipse" />
                    <ion-label>Tab 2</ion-label>
                </ion-tab-button>
                
                <ion-tab-button tab="settings" href="/tabs/settings">
                    <ion-icon :icon="settingsSharp" />
                    <ion-label>Settings</ion-label>
                </ion-tab-button>
                
            </ion-tab-bar>
        </ion-tabs>
    </ion-page>
</template>

<script lang="ts">
import { IonTabBar, IonTabButton, IonTabs, IonLabel, IonIcon, IonPage, IonRouterOutlet } from '@ionic/vue';
import { ellipse, settingsSharp, homeSharp } from 'ionicons/icons';

export default {
    name: 'Tabs',
    components: { IonLabel, IonTabs, IonTabBar, IonTabButton, IonIcon, IonPage, IonRouterOutlet },
    setup() {
        return { ellipse, settingsSharp, homeSharp };
    }
}
</script>
