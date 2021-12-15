<template>
  <v-app>
    <v-app-bar
      app
      color="#43a047"
      dark

    >
      <div class="d-flex align-center">
        <v-img
          alt="Gossamer Logo"
          class="shrink mr-2"
          contain
          src="./assets/logo_app.png"
         
          width="160"
        /> 

       
      </div>

     
      <v-spacer></v-spacer>

     
    </v-app-bar>


    <v-main>
   
      <MainBoard/>
    </v-main>
     <v-footer
    dark
    padless
  >
    <v-card
      class="flex"
      flat
      tile
    >
      <v-card-title class="">
        <strong class="subheading">Панель управления</strong>

        <v-spacer></v-spacer>
      <v-btn
        href="https://disk.yandex.ru/d/eoaonlylef2lEg"
        target="_blank"
        text
      >
        
        <!-- <v-icon>mdi-open-in-new</v-icon> -->
        <v-icon class="mr-2"  >
          mdi-android 
        </v-icon>
        <span class="mr-2">Мобильное приложение</span>
      </v-btn>
       <v-btn
        href="https://disk.yandex.ru/d/UYqtGKNTVe57hg" 
        target="_blank"
        text
      >
        
        <!-- <v-icon>mdi-open-in-new</v-icon> -->
        <v-icon class="mr-2"  >
          mdi-bookshelf  
        </v-icon>
        <span class="mr-2">Лабораторные работы</span>
      </v-btn>

        <v-btn
        @click="dialog = true"
        
        text
      >
        
        <!-- <v-icon>mdi-open-in-new</v-icon> -->
        <v-icon class="mr-2"  >
          mdi-cog 
        </v-icon>
        <span class="mr-2">Настройки</span>
      </v-btn>


      <div>
    <v-row
      justify="center"
    >
      
 
      <v-dialog
        v-model="dialog"
        fullscreen
        hide-overlay
        transition="dialog-bottom-transition"
        scrollable
      >
        <v-card tile>
          <v-toolbar
            flat
            dark
            color="primary"
          >
            <v-btn
              icon
              dark
              @click="dialog = false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-toolbar-title>Настройки</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <v-btn
                dark
                text
                @click="dialog = false; saveLinkApi();"
              >
                Сохранить
              </v-btn>
            </v-toolbar-items>
             
          </v-toolbar>
          
          <v-card-text>
          
            
            <v-list
              three-line
              subheader
            >
              <v-subheader>Основные</v-subheader>
 
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>Точка подключения</v-list-item-title>
                  <v-list-item-subtitle>Укажите адрес сервера (пример: "http://localhost:8080")</v-list-item-subtitle>
                 <v-text-field v-model="linkApi"
                      label="http://"></v-text-field>
                </v-list-item-content>
              </v-list-item>
            </v-list>
            <v-divider></v-divider>
       
          </v-card-text>

          <div style="flex: 1 1 auto;"></div>
        </v-card>
      </v-dialog>

      
       
    </v-row>
  </div>
      </v-card-title>

      

      <v-card-text class="py-2 white--text text-center">
        {{ new Date().getFullYear() }} — <strong>Gossamer</strong>
      </v-card-text>
    </v-card>
  </v-footer>
  </v-app>
  

  
</template>


<script>
import MainBoard from './components/MainBoard';

export default {
  name: 'App',

  components: {
    MainBoard,
  },

   data: () => ({
       
        linkApi: localStorage.getItem('gossamer_link_api'),
        dialog: false,
        dialog2: false,
        dialog3: false,
        notifications: false,
        sound: true,
        widgets: false,
   
      
    }),
    methods: {
      saveLinkApi(){
        alert(this.linkApi);
        localStorage.setItem('gossamer_link_api', this.linkApi);
        //alert( localStorage.getItem('test') ); // 1
      },
      initialize () {
        this.linkApi =" localStorage.getItem('gossamer_link_api')";
         alert(this.linkApi);
      }
    },
};
</script>
