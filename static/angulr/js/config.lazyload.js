// lazyload config

angular.module('app')
    /**
   * jQuery plugin config use ui-jq directive , config the js and css files that required
   * key: function name of the jQuery plugin
   * value: array of the css js file located
   */
  .constant('JQ_CONFIG', {
      easyPieChart:   [   '/greenhome/static/libs/jquery/jquery.easy-pie-chart/dist/jquery.easypiechart.fill.js'],
      sparkline:      [   '/greenhome/static/libs/jquery/jquery.sparkline/dist/jquery.sparkline.retina.js'],
      plot:           [   '/greenhome/static/libs/jquery/flot/jquery.flot.js',
                          '/greenhome/static/libs/jquery/flot/jquery.flot.pie.js', 
                          '/greenhome/static/libs/jquery/flot/jquery.flot.resize.js',
                          '/greenhome/static/libs/jquery/flot.tooltip/js/jquery.flot.tooltip.min.js',
                          '/greenhome/static/libs/jquery/flot.orderbars/js/jquery.flot.orderBars.js',
                          '/greenhome/static/libs/jquery/flot-spline/js/jquery.flot.spline.min.js'],
      moment:         [   '/greenhome/static/libs/jquery/moment/moment.js'],
      screenfull:     [   '/greenhome/static/libs/jquery/screenfull/dist/screenfull.min.js'],
      slimScroll:     [   '/greenhome/static/libs/jquery/slimscroll/jquery.slimscroll.min.js'],
      sortable:       [   '/greenhome/static/libs/jquery/html5sortable/jquery.sortable.js'],
      nestable:       [   '/greenhome/static/libs/jquery/nestable/jquery.nestable.js',
                          '/greenhome/static/libs/jquery/nestable/jquery.nestable.css'],
      filestyle:      [   '/greenhome/static/libs/jquery/bootstrap-filestyle/src/bootstrap-filestyle.js'],
      slider:         [   '/greenhome/static/libs/jquery/bootstrap-slider/bootstrap-slider.js',
                          '/greenhome/static/libs/jquery/bootstrap-slider/bootstrap-slider.css'],
      chosen:         [   '/greenhome/static/libs/jquery/chosen/chosen.jquery.min.js',
                          '/greenhome/static/libs/jquery/chosen/bootstrap-chosen.css'],
      TouchSpin:      [   '/greenhome/static/libs/jquery/bootstrap-touchspin/dist/jquery.bootstrap-touchspin.min.js',
                          '/greenhome/static/libs/jquery/bootstrap-touchspin/dist/jquery.bootstrap-touchspin.min.css'],
      wysiwyg:        [   '/greenhome/static/libs/jquery/bootstrap-wysiwyg/bootstrap-wysiwyg.js',
                          '/greenhome/static/libs/jquery/bootstrap-wysiwyg/external/jquery.hotkeys.js'],
      dataTable:      [   '/greenhome/static/libs/jquery/datatables/media/js/jquery.dataTables.min.js',
                          '/greenhome/static/libs/jquery/plugins/integration/bootstrap/3/dataTables.bootstrap.js',
                          '/greenhome/static/libs/jquery/plugins/integration/bootstrap/3/dataTables.bootstrap.css'],
      vectorMap:      [   '/greenhome/static/libs/jquery/bower-jvectormap/jquery-jvectormap-1.2.2.min.js', 
                          '/greenhome/static/libs/jquery/bower-jvectormap/jquery-jvectormap-world-mill-en.js',
                          '/greenhome/static/libs/jquery/bower-jvectormap/jquery-jvectormap-us-aea-en.js',
                          '/greenhome/static/libs/jquery/bower-jvectormap/jquery-jvectormap.css'],
      footable:       [   '/greenhome/static/libs/jquery/footable/v3/js/footable.min.js',
                          '/greenhome/static/libs/jquery/footable/v3/css/footable.bootstrap.min.css'],
      fullcalendar:   [   '/greenhome/static/libs/jquery/moment/moment.js',
                          '/greenhome/static/libs/jquery/fullcalendar/dist/fullcalendar.min.js',
                          '/greenhome/static/libs/jquery/fullcalendar/dist/fullcalendar.css',
                          '/greenhome/static/libs/jquery/fullcalendar/dist/fullcalendar.theme.css'],
      daterangepicker:[   '/greenhome/static/libs/jquery/moment/moment.js',
                          '/greenhome/static/libs/jquery/bootstrap-daterangepicker/daterangepicker.js',
                          '/greenhome/static/libs/jquery/bootstrap-daterangepicker/daterangepicker-bs3.css'],
      tagsinput:      [   '/greenhome/static/libs/jquery/bootstrap-tagsinput/dist/bootstrap-tagsinput.js',
                          '/greenhome/static/libs/jquery/bootstrap-tagsinput/dist/bootstrap-tagsinput.css']
                      
    }
  )
  .constant('MODULE_CONFIG', [
      {
          name: 'ngGrid',
          files: [
              '/greenhome/static/libs/angular/ng-grid/build/ng-grid.min.js',
              '/greenhome/static/libs/angular/ng-grid/ng-grid.min.css',
              '/greenhome/static/libs/angular/ng-grid/ng-grid.bootstrap.css'
          ]
      },
      {
          name: 'ui.grid',
          files: [
              '/greenhome/static/libs/angular/angular-ui-grid/ui-grid.min.js',
              '/greenhome/static/libs/angular/angular-ui-grid/ui-grid.min.css',
              '/greenhome/static/libs/angular/angular-ui-grid/ui-grid.bootstrap.css'
          ]
      },
      {
          name: 'ui.select',
          files: [
              '/greenhome/static/libs/angular/angular-ui-select/dist/select.min.js',
              '/greenhome/static/libs/angular/angular-ui-select/dist/select.min.css'
          ]
      },
      {
          name:'angularFileUpload',
          files: [
            '/greenhome/static/libs/angular/angular-file-upload/angular-file-upload.js'
          ]
      },
      {
          name:'ui.calendar',
          files: ['/greenhome/static/libs/angular/angular-ui-calendar/src/calendar.js']
      },
      {
          name: 'ngImgCrop',
          files: [
              '/greenhome/static/libs/angular/ngImgCrop/compile/minified/ng-img-crop.js',
              '/greenhome/static/libs/angular/ngImgCrop/compile/minified/ng-img-crop.css'
          ]
      },
      {
          name: 'angularBootstrapNavTree',
          files: [
              '/greenhome/static/libs/angular/angular-bootstrap-nav-tree/dist/abn_tree_directive.js',
              '/greenhome/static/libs/angular/angular-bootstrap-nav-tree/dist/abn_tree.css'
          ]
      },
      {
          name: 'toaster',
          files: [
              '/greenhome/static/libs/angular/angularjs-toaster/toaster.js',
              '/greenhome/static/libs/angular/angularjs-toaster/toaster.css'
          ]
      },
      {
          name: 'textAngular',
          files: [
              '/greenhome/static/libs/angular/textAngular/dist/textAngular-sanitize.min.js',
              '/greenhome/static/libs/angular/textAngular/dist/textAngular.min.js'
          ]
      },
      {
          name: 'vr.directives.slider',
          files: [
              '/greenhome/static/libs/angular/venturocket-angular-slider/build/angular-slider.min.js',
              '/greenhome/static/libs/angular/venturocket-angular-slider/build/angular-slider.css'
          ]
      },
      {
          name: 'com.2fdevs.videogular',
          files: [
              '/greenhome/static/libs/angular/videogular/videogular.min.js'
          ]
      },
      {
          name: 'com.2fdevs.videogular.plugins.controls',
          files: [
              '/greenhome/static/libs/angular/videogular-controls/controls.min.js'
          ]
      },
      {
          name: 'com.2fdevs.videogular.plugins.buffering',
          files: [
              '/greenhome/static/libs/angular/videogular-buffering/buffering.min.js'
          ]
      },
      {
          name: 'com.2fdevs.videogular.plugins.overlayplay',
          files: [
              '/greenhome/static/libs/angular/videogular-overlay-play/overlay-play.min.js'
          ]
      },
      {
          name: 'com.2fdevs.videogular.plugins.poster',
          files: [
              '/greenhome/static/libs/angular/videogular-poster/poster.min.js'
          ]
      },
      {
          name: 'com.2fdevs.videogular.plugins.imaads',
          files: [
              '/greenhome/static/libs/angular/videogular-ima-ads/ima-ads.min.js'
          ]
      },
      {
          name: 'xeditable',
          files: [
              '/greenhome/static/libs/angular/angular-xeditable/dist/js/xeditable.min.js',
              '/greenhome/static/libs/angular/angular-xeditable/dist/css/xeditable.css'
          ]
      },
      {
          name: 'smart-table',
          files: [
              '/greenhome/static/libs/angular/angular-smart-table/dist/smart-table.min.js'
          ]
      },
      {
          name: 'angular-skycons',
          files: [
              '/greenhome/static/libs/angular/angular-skycons/angular-skycons.js'
          ]
      }
    ]
  )
  // oclazyload config
  .config(['$ocLazyLoadProvider', 'MODULE_CONFIG', function($ocLazyLoadProvider, MODULE_CONFIG) {
      // We configure ocLazyLoad to use the lib script.js as the async loader
      $ocLazyLoadProvider.config({
          debug:  false,
          events: true,
          modules: MODULE_CONFIG
      });
  }])
;
