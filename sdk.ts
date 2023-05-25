import { AxiosInstance, AxiosResponse } from "axios";


export interface PayloadDataDD {
    Name: string;
}


class ClientSdk {
    client!: AxiosInstance

}
const clientSdk = new ClientSdk()

export function SetClient(client: AxiosInstance) {
    clientSdk.client = client
}


export async function PostUsers(query: any, data: PayloadDataDD) {
    let res = await clientSdk.client.post<any, AxiosResponse<any, any>, PayloadDataDD>('/users', data, {
        params: query,
    });
    return res.data;
}


export async function PostProductCreate(query: any, data: PayloadDataDD) {
    let res = await clientSdk.client.post<any, AxiosResponse<any, any>, PayloadDataDD>('/product/create', data, {
        params: query,
    });
    return res.data;
}
