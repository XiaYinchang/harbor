
import {Injectable} from "@angular/core";
import {Http} from "@angular/http";
import { HTTP_JSON_OPTIONS } from "./shared/shared.utils";

@Injectable()
export class SkinableConfig {
    customSkinData: {[key: string]: any};
    constructor(private http: Http) {}

    public getCustomFile(): Promise<any> {
       return this.http.get('setting.json', HTTP_JSON_OPTIONS)
           .toPromise()
           .then(response => { this.customSkinData = response.json(); return this.customSkinData; })
           .catch(error => {
               console.error('custom skin json file load failed');
           });
    }

    public getSkinConfig() {
        return this.customSkinData;
    }

    public getProject() {
        if (this.customSkinData) {
            return this.customSkinData.project;
        } else {
            return null;
        }
    }
}
