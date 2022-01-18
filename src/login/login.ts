import { Component, NgZone } from '@angular/core';
import {Router} from '@angular/router';
import { time } from 'console';
import { HttpService } from '../services/http.service';
     
@Component({
    selector: 'log-app',
    templateUrl: './login.html',
    styleUrls: ['./login.css'], 
    providers: [HttpService]
})
export class LoginComponent { 
    constructor(
        private router: Router,
        private httpService: HttpService,
        private ngZone: NgZone){}

    typePasswordFild = "password";
    textButtomShowPassword = "show-password";
    flPassword : boolean = true;
    flLogin : boolean = true;
    flCorrectInput : boolean = true;
    login : string = '';
    password : string = '';
    text : Text | undefined;
    texts: Text[]=[];
    LoginOk : boolean;

    ngOnInit(){
        this.LoginOk = true;
        console.log(window.location.href);
        this.httpService.getData("http://localhost:8080/angular").subscribe((data:any) => this.texts = data);
    }


    enter(){
        this.flCorrectInput = true;
        
        //console.log(this.login);
        if (this.login === '') {
            this.flLogin = false;
            this.flCorrectInput = false;
        }
        else{
            this.flLogin = true;
        }
        if (this.password === '') {
            this.flPassword = false;
            this.flCorrectInput = false;
        }
        else{
            this.flPassword = true;
        }


        if (this.flCorrectInput === true) {
            const data = {login: this.login, password : this.password};

            this.ngZone.run( _ => {
                this.httpService.postData("http://localhost:8080/login", this.login + " " + this.password).subscribe((value:any) => {
                    this.LoginOk = value.FlLogin;
                    console.log("вот", this.LoginOk)

                    if(this.LoginOk == true)
                    {
                        this.router.navigate(['admin']);
                    }
                    else
                    {
                        console.log("ошибка")
                    }
                        });
                    })
            
           // this.checkPassword()
        }
       // console.log(this.password);
       console.log("вот", this.LoginOk)
    }


    showPassword() {
        if (this.typePasswordFild === "password") {
            this.typePasswordFild = "text";
        }
        else if (this.typePasswordFild === "text") {
            this.typePasswordFild = "password";
        }

        if (this.textButtomShowPassword === "show-password") {
            this.textButtomShowPassword = "hide-password";
        }
        else if (this.textButtomShowPassword === "hide-password") {
            this.textButtomShowPassword = "show-password";
        }
    }

    
}