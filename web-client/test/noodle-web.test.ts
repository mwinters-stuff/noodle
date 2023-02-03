import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';

import { NoodleWeb } from '../src/noodle-web.js';

describe('NoodleWeb', () => {
  let element: NoodleWeb;
  beforeEach(async () => {
    element = await fixture(html`<noodle-web></noodle-web>`);
  });

  it('renders a h1', () => {
    const h1 = element.shadowRoot!.querySelector('h1')!;
    expect(h1).to.exist;
    expect(h1.textContent).to.equal('My app');
  });

  it('passes the a11y audit', async () => {
    await expect(element).shadowDom.to.be.accessible();
  });
});
