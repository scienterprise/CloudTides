import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { environment } from '@tide-environments/environment';
import { VCD_URL_PATH } from '@tide-config/path';
import { LoginService } from '../login/login.service';
import toFixed from 'accounting-js/lib/toFixed.js';

@Injectable()
export class VendorService {

  constructor(
    private readonly http: HttpClient,
    private readonly loginService: LoginService,
  ) {
  }

  private prefix = `${environment.apiPrefix}/computeResource`;

  async getList() {
    const list = await this.http.get<Item[]>(environment.apiPrefix + VCD_URL_PATH, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise();
    const vendors: Item[] = [];
    for (const vendor of list) {
      const vendorItem: Item = {
        id: vendor.id,
        name: vendor.name,
        version: vendor.version,
        type: vendor.type,
        url: vendor.url,
      };
      vendors.push(vendorItem);
    }
    return vendors;
  }

  addItem(payload: ItemPayload) {
    const body = {
      ...payload,
      policy: 0,
    };
    return this.http.post<any>(environment.apiPrefix + VCD_URL_PATH, body, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise().then(() => {
      return Promise.resolve();
    }, (errResp) => {
      return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
    });
  }

  editItem(id: string, payload: ItemPayload) {
    return this.http.put<ItemDTO>(`${this.prefix}/${id}`, payload).pipe(
      map(mapItem),
    );
  }

  async removeItem(id: string) {
    await this.http.delete<any>(environment.apiPrefix + `/resource/vcd/` + id, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      () => {
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
  }

  async contributeResource(id: string): Promise<ContributeResp> {
    let response = null;
    await this.http.put(environment.apiPrefix + `/resource/contribute/${id}`, null, {
      headers: {
      Authorization: `Bearer ${this.loginService.token}`,
    }, }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }

  async activateResource(id: string): Promise<ActivateResp> {
    let response = null;
    await this.http.put<ActivateResp>(environment.apiPrefix + `/resource/activate/${id}`, null, {
      headers: {
        Authorization: `Bearer ${this.loginService.token}`,
      },
    }).toPromise().then(
      (resp) => {
        response = resp;
        return Promise.resolve();
      }, (errResp) => {
        return Promise.reject(`${errResp.message}`);
      },
    );
    return response;
  }
}

interface ItemDTO {
  id: number;
  name: string;
  url: string;
  type: string;
  version: number;
}

interface ContributeResp {
  message: string;
  contributed: boolean;
}

interface ActivateResp {
  message: string;
  activated: boolean;
}

function mapList(raw: ItemDTO[]): Item[] {
  return raw.map(mapItem);
}

function mapItem(raw: ItemDTO): Item {
  return raw;
}

// UI
export interface ItemPayload {
  datacenter: string;
  name: string;
  org: string;
  network: string;
  catalog: string;
  username: string,
  password: string,
}

export type Item = ItemDTO;
