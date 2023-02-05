import { html, css } from 'lit';
// eslint-disable-next-line import/no-extraneous-dependencies
import { Commands, Context, Router } from '@vaadin/router';
import { customElement } from 'lit/decorators.js';

import './noodle-login.js';
import './noodle-dash.js';
import { NoodleAuthenticatedBase } from './noodle-authenticated-base.js';

@customElement('noodle-web')
export class NoodleWeb extends NoodleAuthenticatedBase {
  static styles = css`
    :host {
      display: block;
      border-width: 0;
      width: 100%;
      height: 100%;
    }
    main {
      width: 100%;
      height: 100%;
      border-width: 0;
      box-sizing: border-box;
    }
  `;

  activeRoute: string = '';

  params: string = '';

  query: string = '';

  data: string = '';

  static async logout(context: Context, commands: Commands) {
    const response =
      await NoodleAuthenticatedBase.getAuthenticatedApi().auth.logoutList();
    if (response.error != null) {
      console.error(response.error);
      // this._errorMessage = value.error.message ?? 'Unknown Error';
    }
    document.cookie =
      'noodle-auth=; expires=Thu, 01 Jan 1970 00:00:00 GMT; Secure';

    return commands.redirect('/login'); // pass to the next route in the list
  }

  firstUpdated() {
    super.firstUpdated();
    const router = new Router(this.shadowRoot!.querySelector('main'));
    router.setRoutes([
      { path: '/dash', component: 'noodle-dash' },
      { path: '/login', component: 'noodle-login' },
      { path: '/logout', action: NoodleWeb.logout },
      {
        path: '(.*)',
        redirect: '/dash',
      },
    ]);
    if (!NoodleAuthenticatedBase.IsAuthenticated()) {
      Router.go(`/login`);
    }
  }

  render() {
    return html` <main></main> `;
  }
}
