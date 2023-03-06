
import '@shoelace-style/shoelace/dist/components/icon/icon.js'
import '@shoelace-style/shoelace/dist/components/alert/alert.js'
import { ResponseError } from '../api/index.js';

export class Functions {
  public static DarkColor = "rgb(22,33,37)";

  public static LightColor = "rgb(250,250,250)";

  public static modifyColor(color: string): string {
    if (color === "dark") {
      return Functions.DarkColor;
    }
    if (color === "light") {
      return Functions.LightColor;
    }
    if (color !== "") {
      return color;
    }

    return Functions.DarkColor

  }

  public static showWebResponseError(reason: ResponseError) {
    reason.response.json().then((value: any) => {
      Functions.showError(value.message);
    });
  }

  public static showError(error: string) {
    Functions.notify(`<strong>Error</strong><br />${error}`, "danger", "exclamation-octagon")
  }


  // Always escape HTML for text arguments!
  public static escapeHtml(html: string): string {
    const div = document.createElement('div');
    div.textContent = html;
    return div.innerHTML;
  }

  // Custom function to emit toast notifications
  public static notify(message: string, variant: string = 'primary', icon: string = 'info-circle', durationx: number = 3000) {
    const alert = Object.assign(document.createElement('sl-alert'), {
      variant,
      closable: true,
      duration: durationx === 0 ? null : durationx,
      innerHTML: `
          <sl-icon name="${icon}" slot="icon"></sl-icon>
          ${message}
        `
    });

    document.body.append(alert);
    return alert.toast();
  }
}