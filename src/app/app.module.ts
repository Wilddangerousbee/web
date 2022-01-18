import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }   from '@angular/forms';
import {Routes, RouterModule} from '@angular/router';
import { HttpClientModule } from '@angular/common/http'

import { AppComponent }   from '../companent/app.component';
import { LoginComponent } from "../login/login";
import { AdminComponent } from "../admin/admin";
import { CardComponent } from "../card/card";
import { ResumeComponent } from "../resume/resume"

// определение дочерних маршрутов
const resumeRoutes: Routes = [
    { path: 'resume', component: ResumeComponent},
];

const appRoutes: Routes =[
    {path: '', component: LoginComponent},
    {path: 'admin', component: AdminComponent},
    {path: 'resume/:resumeId', component: ResumeComponent}
]

@NgModule({
    imports:      [ BrowserModule, FormsModule, RouterModule.forRoot(appRoutes), HttpClientModule],
    declarations: [ AppComponent , LoginComponent, AdminComponent, CardComponent, ResumeComponent],
    bootstrap:    [ AppComponent ]
})
export class AppModule { }