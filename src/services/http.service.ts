import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
  
@Injectable()
export class HttpService{
    urlUsers : string = "assets/users.json";
    urlAuf : string = "http://localhost:3000";
  
    constructor(private http: HttpClient){ }
      
    getData(url : string){
        return this.http.get(url);
    }

    postData(url : string, data:any){
        console.log(window.location.href);
        return this.http.post(url, data);
    }

    updateData(url : string, data : any) {
        return this.http.put(url, data);
    }
}

