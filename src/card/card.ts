import { Component , Input } from '@angular/core';
import { HttpService} from '../services/http.service';
import { AdminComponent } from '../admin/admin';
import { User } from '../datamodel/user';
     
@Component({
    selector: 'card-companent',
    templateUrl: './card.html',
    styleUrls: ['./card.css']
})

export class CardComponent { 
    @Input() name: string;
}