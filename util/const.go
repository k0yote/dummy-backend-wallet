package util

const mailTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head data-id="__react-email-head">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  </head>
  <body
    data-id="__react-email-body"
    style="
      margin-left: auto;
      margin-right: auto;
      overflow: auto;
      background-color: rgb(243, 242, 247);
      font-family: Inter, ui-sans-serif, system-ui, -apple-system,
        BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans,
        sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol,
        Noto Color Emoji;
    "
  >
    <table
      align="center"
      width="100%"
      data-id="__react-email-container"
      role="presentation"
      cellspacing="0"
      cellpadding="0"
      border="0"
      style="
        max-width: 520px;
        margin-left: auto;
        margin-right: auto;
        margin-top: 2.5rem;
        background-color: rgb(255, 255, 255);
        padding-left: 4rem;
        padding-right: 4rem;
        padding-top: 2rem;
        padding-bottom: 2rem;
        text-align: center;
        font-weight: 400;
        color: rgb(22, 11, 69);
      "
    >
      <tbody>
        <tr style="width: 100%">
          <td>
            <p
              data-id="react-email-text"
              style="
                font-size: 12px;
                line-height: 20px;
                margin: 16px 0;
                margin-bottom: 0px;
                margin-top: 0.5rem;
                text-align: center;
                letter-spacing: -0.0025em;
                color: rgb(173, 169, 188);
              "
            >
              Dummy Auth Demo
            </p>
            <h1
              data-id="react-email-heading"
              style="
                margin-top: 0.25rem;
                font-size: 18px;
                line-height: 24px;
                letter-spacing: 0.01%;
                font-weight: 400;
              "
            >
              Log in to
              <strong style="font-weight: 700">Dummy Auth Demo</strong>
            </h1>
            <hr
              data-id="react-email-hr"
              style="
                width: 100%;
                border: none;
                border-top: 1px solid #eaeaea;
                margin-top: 1rem;
                border-width: 1px;
                border-color: rgb(232, 230, 239);
              "
            />
            <p
              data-id="react-email-text"
              style="
                font-size: 13px;
                line-height: 20px;
                margin: 16px 0;
                margin-bottom: 0.5rem;
                margin-top: 1.5rem;
                letter-spacing: -0.0025em;
              "
            >
              Your code is
            </p>
            <table
              align="center"
              width="100%"
              data-id="react-email-section"
              border="0"
              cellpadding="0"
              cellspacing="0"
              role="presentation"
              style="
                margin-left: auto;
                margin-right: auto;
                margin-top: 0px;
                width: fit-content;
                border-radius: 12px;
                border-width: 1px;
                border-style: solid;
                border-color: rgb(232, 230, 239);
                padding-left: 1.5rem;
                padding-right: 1.5rem;
                padding-top: 1rem;
                padding-bottom: 1rem;
              "
            >
              <tbody>
                <tr>
                  <td>
                    <p
                      data-id="react-email-text"
                      style="
                        font-size: 32px;
                        line-height: 24px;
                        margin: 16px 0;
                        font-weight: 500;
                      "
                    >
                      %v
                    </p>
                  </td>
                </tr>
              </tbody>
            </table>
            <p
              data-id="react-email-text"
              style="
                font-size: 13px;
                line-height: 20px;
                margin: 16px 0;
                margin-top: 0.75rem;
                letter-spacing: -0.0025em;
              "
            >
              This code expires in 10 minutes. Do not share this code with
              anyone.
            </p>
            <hr
              data-id="react-email-hr"
              style="
                width: 100%;
                border: none;
                border-top: 1px solid #eaeaea;
                margin-top: 2rem;
                border-width: 1px;
                border-color: rgb(232, 230, 239);
              "
            />
            <p
              data-id="react-email-text"
              style="
                font-size: 13px;
                line-height: 20px;
                margin: 16px 0;
                margin-bottom: 0px;
                text-align: left;
                letter-spacing: -0.0025em;
              "
            >
              Wasn't you? Please email
              <a href="mailto:simplewallettest6@gmail.com">simplewallettest6@gmail.com</a> with any
              questions
            </p>
            <p
              data-id="react-email-text"
              style="
                font-size: 11px;
                line-height: 24px;
                margin: 16px 0;
                margin-top: 0px;
                margin-bottom: 0px;
                text-align: left;
                color: rgb(119, 113, 152);
              "
            >
              Platform: ?????
            </p>
          </td>
        </tr>
      </tbody>
    </table>
    <p
      data-id="react-email-text"
      style="
        font-size: 13px;
        line-height: 20px;
        margin: 16px 0;
        margin-left: auto;
        margin-right: auto;
        display: flex;
        width: fit-content;
        text-align: center;
        letter-spacing: -0.0025em;
        color: rgb(173, 169, 188);
      "
    >
      Powered by Dummy
    </p>
    <img
      src="http://url5441.privy.io/wf/open?upn=3DBtixqU2mt4KOrXFsl7KBE0lM7k9i7RHHgceywEyVQZMljaDaabbAEFiQ9SdB6hcplLoreYrdwquJP2zcn9YUG4fiOEXP-2BBbjU-2B7EUCzioC0aL-2FXgx5ew9WH6u-2FOouFTgsdSFBO4cZTKRo1dvk-2FsVjRcJuMsW3Poy51xL3NrrEpAx-2BWXvEKARUO26mCC43J6wxfJ7cQKZPNxcRzPXAry9Ply7kH9L15Rg9fE30m19flGb7ZOW4QSS9V-2FGvTENhETi8EsXrAhAiapj7kVk8kdY2nS3LpiR1m9Vh36gHDYtRVWANA-2BOyvl-2B-ulktYLtRkEPbcJ8mzini7sBJfgMeRnQRRWw-3D-3D"
      alt=""
      width="1"
      height="1"
      border="0"
      style="
        height: 1px !important;
        width: 1px !important;
        border-width: 0 !important;
        margin-top: 0 !important;
        margin-bottom: 0 !important;
        margin-right: 0 !important;
        margin-left: 0 !important;
        padding-top: 0 !important;
        padding-bottom: 0 !important;
        padding-right: 0 !important;
        padding-left: 0 !important;
      "
    />
  </body>
</html>
`
