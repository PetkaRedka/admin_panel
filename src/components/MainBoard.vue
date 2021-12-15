
<template>

  <v-data-table
    
    :headers="headers"
    :items="studentList"
    sort-by="calories"
    class="elevation-1"
    hide-default-footer
  >

    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title>Список студентов</v-toolbar-title>
    
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="500px">
          <template v-slot:activator="{ on, attrs }">
            
            <v-btn color="primary" dark class="mb-2" v-bind="attrs" v-on="on">
               <v-icon class="mr-2"  >
             mdi-account-plus
            </v-icon>
              Добавить
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                <v-col cols="12" >
                    <v-text-field v-model="editedItem.given_name"
                      label="Имя студента"></v-text-field>
                  </v-col>
                   </v-row>
                    <v-row>
                  <v-col cols="12">
                    <v-text-field v-model="editedItem.surname"
                      label="Фамилия студента"></v-text-field>
                  </v-col>
                   </v-row>
                    <v-row>
                   <v-col cols="12">
                    <v-text-field
                      v-model="editedItem.class"
                      label="Номер группы"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="6" md="4">
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="close">
                Отмена
              </v-btn>
              <v-btn color="blue darken-1" text @click="saveBtn">
                Сохранить
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="dialogDelete" max-width="500px">
          <v-card>
            <v-card-title class="center">Вы хотите удалить студента?</v-card-title> 
            <v-card-text class="text-md">Удаление студента приведет к удалению его рабочего стола</v-card-text> 
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="closeDelete">Нет</v-btn>
              <v-btn color="blue darken-1" text @click="deleteItemConfirm">Да</v-btn>
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>

     <template v-slot:item.link="{ item }">
      <v-chip
        :color="getColor(item.link)"
        dark
      >
        {{ item.link }}
      </v-chip>
    </template>
   
    <template v-slot:item.actions="{ item }">
      <!--<v-icon small class="mr-2" @click="editItem(item)">
        mdi-pencil
      </v-icon>-->
       <v-icon class="mr-2"  @click="qr_dialog = true; generateQr(item); ">
        mdi-qrcode-scan 
      </v-icon>
      <v-icon class="mr-2"  @click="progress_dialog = true;  runItem(item);">
        mdi-play-circle 
      </v-icon>
      <v-icon class="mr-2"   @click="deleteItem(item)">
        mdi-close-circle
      </v-icon>
 
  <div class="text-center">


 <v-dialog
      v-model="qr_dialog"
     
     
      width="500"
    >
      <v-card
        color="dark"
        dark
      >
        <v-card-text>
         Наведите камеру телефона
         <v-img
        
    
        :src=qr_dialog_link
></v-img>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-dialog
      v-model="progress_dialog"
      hide-overlay
      persistent
      width="300"
    >
      <v-card
        color="primary"
        dark
      >
        <v-card-text>
          Рабочий стол создается...
          <v-progress-linear
            indeterminate
            color="white"
            class="mb-0"
          ></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
 
    </template>
   
    <template v-slot:no-data>
      <v-btn
        color="primary" @click="initialize">
        Сбросить
      </v-btn>
    </template>
  </v-data-table>
  
</template>

<script>
  export default {
    data: () => ({
        qr_dialog: false,
        qr_dialog_link: "http://qrcoder.ru/code/?vnc%3A%2F%2Fvuc%40vuc.xxx.engineer%3A5901&10&0",
        progress_dialog: false,
        dialog: false,
        dialogDelete: false,
        host: localStorage.getItem('gossamer_link_api'),//"https://vuc.xxx.engineer",
        headers: [
           {
            text: 'Имя студента',
            align: 'start',
            sortable: true,
            value: 'given_name',
        },
        {
            text: 'Фамилия студента',
            align: 'start',
            sortable: true,
            value: 'surname',
        },
        { text: 'Номер группы', value: 'class' },
        { text: 'Ссылка для VNC', value: 'link' },
        // { text: 'Статус', value: 'status' },
        { text: 'Сессия (id)', value: 'session', sortable: false, },
        { text: 'Действия', value: 'actions', sortable: false },
        ],
        studentList: [],
        editedIndex: -1,
        editedItem: {
        given_name: '',
        surname: '',
        class: '',
        session: '',
        link: 'Рабочий стол не создан',
        },
        // defaultItem: {
        // given_name: 'Саша',
        // surname: 'Морозов',
        // class: 'КТСО-02-17',
        // session: '',
        // },
}),

    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'Добавление студента' : 'Редактирование'
      },
    },

    watch: {
      dialog (val) {
        val || this.close()
      },
      dialogDelete (val) {
        val || this.closeDelete()
      },
    },

    created () {
      this.initialize()
    },

    methods: {
      generateQr(item){
         if (item.link == "Рабочий стол не создан" || item.link == "") {
           this.qr_dialog = false;
           alert("Рабочий стол не создан!");
           return;
         }  
        var link = item.link;
        link = link.substr(6);
        this.qr_dialog_link = "http://qrcoder.ru/code/?vnc%3A%2F%2F"+link+"&10&0" ;
      
      },
      getColor (link) {
        if (link == "Рабочий стол не создан" || link == "") return 'grey'
        else return 'green'
      },
      initialize () {
        this.studentList = [
        ]

        //var json_response = '{"connections":[{"surname": "Frozen Yogurt4", "class": 8080, "port": 8089, "status": "Online", "session": "asdas21e909"}]}';      
        // var response = JSON.parse(json_response);
        // response.connections.forEach(connection => {
        //    this.studentList.push({
        //           surname: connection.surname,
        //           class: connection.class,
        //           port: connection.port,
        //           status: connection.status,
        //           session: connection.session,
        //         })
 
        // });

        //localStorage.setItem('test', 1);
        //alert( localStorage.getItem('test') ); // 1
        fetch(this.host + "/admin/users", {
            method: "GET",
            headers: {
              "Content-Type": "application/json"
            },
          })
        .then(async response => {
        
       // alert(response.json())
        const data = await response.json();

        //alert(JSON.stringify(data));
        if (!response.ok) {
            // get error message from body or default to response statusText
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

      
        data.users.forEach(user => {
          var link = "Рабочий стол не создан";
          if  (user.session != null) {
            link = user.session.connection_url;
          }

           this.studentList.push({
                  given_name: user.given_name,
                  surname: user.surname,
                  class: user.class,
                  port: "",
                  status: "",
                  link: link,
                  session: user.id,
                })
         });
        })
        .catch(error => {
          this.errorMessage = error;
          alert("Не удалось получить список студентов!\nПроверьте правильность настройки точки подключения");
          console.error("There was an error!", error);
        });

      },

      editItem (item) {
        this.editedIndex = this.studentList.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },

     runItem (item) {
       if(item.link != "Рабочий стол не создан" && item.link != "") {
         alert("Рабочий стол уже создан!\nТочка подключения: " + item.link);
         this.progress_dialog = false;
         return;
       }

       
        fetch(this.host + "/session", {
          method: "POST",
           headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            surname: item.surname, 
            class: item.class
            })
        }).then(async response => {
          const data = await response.json();
       
          if (!response.ok) {    
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
          }

          this.editedItem.link = data['connection_url'];
          
          location.reload();
          // this.save();
          
       
        })
        .catch(error => {
          this.errorMessage = error;
         // alert(error);
          alert("Не удалось создать рабочий стол");
          location.reload();
          console.error("There was an error!", error);
        }
        );


      },

      deleteItem (item) {
        this.editedIndex = this.studentList.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialogDelete = true
      },

      deleteItemConfirm () {
        var session = this.studentList[this.editedIndex].session
        //alert("Удаление студента с номером сессии " + session)
        this.deleteStudent(session)
        this.studentList.splice(this.editedIndex, 1)
        this.closeDelete()
        //location.reload();
      },

      close () {
        this.dialog = false
        this.$nextTick(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        })
        
      },

      closeDelete () {
        this.dialogDelete = false
        this.$nextTick(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        })
      },
      deleteStudent (session) {
        fetch(this.host + "/admin/users/"+session, {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json"
          }
        }).then(async response => {
          const data = await response.json();
       
          if (!response.ok) {    
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
          }

         //alert(data["message"]);
         this.save();
        })
        .catch(error => {
          this.errorMessage = error;
          // alert("Ошибка при удалении студента");
          //location.reload();
          console.error("There was an error!", error);
        }
        );

       
      
      },

      addStudent () {
 
        fetch(this.host + "/admin/users", {
          method: "POST",
           headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            given_name: this.editedItem.given_name,
            surname: this.editedItem.surname, 
            class: this.editedItem.class
            })
        }).then(async response => {
          const data = await response.json();
       
          if (!response.ok) {    
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
          }

          this.editedItem.session = data['user']['id'];
          this.editedItem.link = "Рабочий стол не создан";
          this.save();
          
       
        })
        .catch(error => {
          this.errorMessage = error;
         // alert(error);
          alert("Не удалось добавить пользователя");
          location.reload();
          console.error("There was an error!", error);
        }
        );

       
      },
      editStudent (session) {
         fetch(this.host, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            session: session,
            surname: this.editedItem.surname, 
            class: this.editItem.class})
        }).then(async response => {
          const data = await response.json();
       
          if (!response.ok) {    
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
          }
         

      
         this.save();
        })
        .catch(error => {
          this.errorMessage = error;
          alert(error);
          console.error("There was an error!", error);
        }
        );
       
      },
      saveBtn () {
        //-1 - edit, other - add
        // if (this.editedIndex == -1) {
        //   //alert("Добавляю нового студента")
        //   this.addStudent()
         
        //   //this.save();
        // } else {
        //   var session = this.studentList[this.editedIndex].session
        //   alert("Редактирую студента с номером сессии " + session)
        //   this.editStudent(session)
        // }     

        this.addStudent()
        
        //location.reload();

      },
      save () {
        if (this.editedIndex > -1) {
          Object.assign(this.studentList[this.editedIndex], this.editedItem)
        } else {
          this.studentList.push(this.editedItem)
        }
        this.close()
      },
    },
  }

</script>