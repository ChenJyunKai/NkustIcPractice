import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class MessageService {
  messages: string[] = [];

  addInService(message: string) {
    this.messages.push(message);
  }

  clearInService() {
    this.messages = []
  }
}