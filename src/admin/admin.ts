import { Component, OnInit, Input, NgZone} from '@angular/core';
import { HttpService} from '../services/http.service';
import { User, SolatPanel } from '../datamodel/user';    

import { CardComponent } from "../card/card";
import { ITS_JUST_ANGULAR } from '@angular/core/src/r3_symbols';

@Component({
    selector: 'admin-companent',
    templateUrl: './admin.html',
    styleUrls: ['./admin.css'],
    providers: [HttpService]
})

export class AdminComponent implements OnInit{ 
    //user: User = new User();
    solatPanels: SolatPanel[]=[];
    tableMode: boolean = true;

    users: User[]=[];

    Sqiare : string
    Mail : string  
    user : User | undefined;
    text : Text | undefined;

    constructor(private httpService: HttpService,
            private ngZone: NgZone){}

    get(solatPanel:SolatPanel)
    {    
        sessionStorage.setItem('userName.com', (solatPanel.Name));
        sessionStorage.setItem('userInfo.com', ("Площадь панели: " + solatPanel.Square.toString() +"\n"+ "Количество панелей: " + solatPanel.Quantity.toString()+"\n"+ "Коэффициент полезности панели: " + solatPanel.UtilityCoefficient.toString()));
    }

    ngOnInit(){
        this.ngZone.run(_ => {
            this.httpService.getData("http://localhost:8080/getCur").subscribe((value:any) => {
            this.solatPanels = value;
            console.log(this.solatPanels);
        })
    })
        console.log(window.location.href);
        this.httpService.getData(this.httpService.urlUsers).subscribe((data:any) => this.users = data["userList"]);
    }


    // получаем данные через сервис
    loadProducts() {
        this.httpService.getData(this.httpService.urlUsers).subscribe((data:any) => this.users = data["userList"]);
    }

    // сохранение данных
    save() {
        if (this.user.id = null) {
            this.httpService.postData(this.httpService.urlUsers, this.user).subscribe((data: User) => this.users.push(data))
        } else {
            this.httpService.updateData(this.httpService.urlUsers, this.user).subscribe((data: User) => this.users.push(data))
        }
    }

    editProduct() {

    }

    cancel() {
        this.user = new User();
        this.tableMode = true;
    }

    delete() {
    }

    add() {
        this.cancel();
        this.tableMode = false;
    }

    enter() {
        this.ngZone.run(_ => {
            this.httpService.postData("http://localhost:8080/config", this.Sqiare + " " + this.Mail).subscribe((value:any) => {
            this.solatPanels = value;
            sessionStorage.setItem('config.com', (this.solatPanels.toString()));
            console.log(this.solatPanels);
        })
    })
    }
}