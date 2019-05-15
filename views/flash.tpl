{{if .flash}}
  {{if .flash.success}}
    <script>
    Swal.fire({
      position: 'top-end',
      type: 'success',
      title: '{{.flash.success}}',
      showConfirmButton: false,
      timer: 1500
    })
    </script>
  {{else if .flash.warning}}
    <script>
      swal("Notice!", "{{.flash.warning}}", "warning");
    </script>
  {{else if .flash.error}}
    <script>
    Swal.fire({
      position: 'top-end',
      type: 'error',
      title: 'Oops...',
      text: '{{.flash.error}}',
      footer: '<a href>Why do I have this issue?</a>'
    })
    </script>
  {{else if .flash.notice}}
      <script>
      Swal.fire({
        title: 'Almost done',
        type: 'info',
        width: 800,
        html:
        'You only need to install the Allstats agent on your server! '+
        '<br><br>Use following command to install the agent on your server, then go to Manage Servers and wait until the initial data has been collected.'+
        '<br><br><code class="important">{{.flash.notice}}</code>',
        background: '#fff url(/static/img/banner-overlay1.png)',

      })
      </script>
   {{end}}
{{end}}
