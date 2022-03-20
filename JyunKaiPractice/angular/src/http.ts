import { HttpClient } from '@angular/common/http';

export class http{
    constructor(
        private http: HttpClient) { 
            this.http.get("http://localhost:8080/api/select")
                .subscribe((res) => {
                console.log(res);
            });
        }
}