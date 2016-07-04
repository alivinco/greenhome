'use strict';

/**
 * Config for the router
 */
angular.module('app')
  .run(
    [          '$rootScope', '$state', '$stateParams',
      function ($rootScope,   $state,   $stateParams) {
          $rootScope.$state = $state;
          $rootScope.$stateParams = $stateParams;        
      }
    ]
  )
  .config(
    [          '$stateProvider', '$urlRouterProvider', 'JQ_CONFIG', 'MODULE_CONFIG', 
      function ($stateProvider,   $urlRouterProvider, JQ_CONFIG, MODULE_CONFIG) {
          var layout = "/greenhome/static/angulr/tpl/app.html";
          if(window.location.href.indexOf("material") > 0){
            layout = "/greenhome/static/angulr/tpl/blocks/material.layout.html";
            $urlRouterProvider
              .otherwise('/app/dashboard-v3');
          }else{
            $urlRouterProvider
              .otherwise('/app/dashboard-v1');
          }
          
          $stateProvider
              .state('app', {
                  abstract: true,
                  url: '/app',
                  templateUrl: layout
              })
              .state('app.dashboard-v1', {
                  url: '/dashboard-v1',
                  templateUrl: '/greenhome/static/angulr/tpl/app_dashboard_v1.html',
                  resolve: load(['/greenhome/static/angulr/js/controllers/chart.js'])
              })
              .state('app.dashboard-v2', {
                  url: '/dashboard-v2',
                  templateUrl: '/greenhome/static/angulr/tpl/app_dashboard_v2.html',
                  resolve: load(['/greenhome/static/angulr/js/controllers/chart.js'])
              })
              .state('app.dashboard-v3', {
                  url: '/dashboard-v3',
                  templateUrl: '/greenhome/static/angulr/tpl/app_dashboard_v3.html',
                  resolve: load(['/greenhome/static/angulr/js/controllers/chart.js'])
              })
              .state('app.ui', {
                  url: '/ui',
                  template: '<div ui-view class="fade-in-up"></div>'
              })
              .state('app.ui.buttons', {
                  url: '/buttons',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_buttons.html'
              })
              .state('app.ui.icons', {
                  url: '/icons',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_icons.html'
              })
              .state('app.ui.grid', {
                  url: '/grid',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_grid.html'
              })
              .state('app.ui.widgets', {
                  url: '/widgets',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_widgets.html'
              })          
              .state('app.ui.bootstrap', {
                  url: '/bootstrap',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_bootstrap.html'
              })
              .state('app.ui.sortable', {
                  url: '/sortable',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_sortable.html'
              })
              .state('app.ui.scroll', {
                  url: '/scroll',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_scroll.html',
                  resolve: load('/greenhome/static/angulr/js/controllers/scroll.js')
              })
              .state('app.ui.portlet', {
                  url: '/portlet',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_portlet.html'
              })
              .state('app.ui.timeline', {
                  url: '/timeline',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_timeline.html'
              })
              .state('app.ui.tree', {
                  url: '/tree',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_tree.html',
                  resolve: load(['angularBootstrapNavTree', '/greenhome/static/angulr/js/controllers/tree.js'])
              })
              .state('app.ui.toaster', {
                  url: '/toaster',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_toaster.html',
                  resolve: load(['toaster', '/greenhome/static/angulr/js/controllers/toaster.js'])
              })
              .state('app.ui.jvectormap', {
                  url: '/jvectormap',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_jvectormap.html',
                  resolve: load('/greenhome/static/angulr/js/controllers/vectormap.js')
              })
              .state('app.ui.googlemap', {
                  url: '/googlemap',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_googlemap.html',
                  resolve: load(['/greenhome/static/angulr/js/app/map/load-google-maps.js', '/greenhome/static/angulr/js/app/map/ui-map.js', '/greenhome/static/angulr/js/app/map/map.js'], function(){ return loadGoogleMaps(); })
              })
              .state('app.ui.vieweditor', {
                  url: '/vieweditor',
                  templateUrl: '/greenhome/static/angulr/tpl/greenhome/view_editor.html',
                  controller: 'XeditableCtrl',
                  resolve: load(['xeditable','/greenhome/static/angulr/js/controllers/xeditable.js'])
              })
              .state('app.chart', {
                  url: '/chart',
                  templateUrl: '/greenhome/static/angulr/tpl/ui_chart.html',
                  resolve: load('/greenhome/static/angulr/js/controllers/chart.js')
              })
              // table
              .state('app.table', {
                  url: '/table',
                  template: '<div ui-view></div>'
              })
              .state('app.table.static', {
                  url: '/static',
                  templateUrl: '/greenhome/static/angulr/tpl/table_static.html'
              })
              .state('app.table.datatable', {
                  url: '/datatable',
                  templateUrl: '/greenhome/static/angulr/tpl/table_datatable.html'
              })
              .state('app.table.footable', {
                  url: '/footable',
                  templateUrl: '/greenhome/static/angulr/tpl/table_footable.html'
              })
              .state('app.table.grid', {
                  url: '/grid',
                  templateUrl: '/greenhome/static/angulr/tpl/table_grid.html',
                  resolve: load(['ngGrid','/greenhome/static/angulr/js/controllers/grid.js'])
              })
              .state('app.table.uigrid', {
                  url: '/uigrid',
                  templateUrl: '/greenhome/static/angulr/tpl/table_uigrid.html',
                  resolve: load(['ui.grid','/greenhome/static/angulr/js/controllers/uigrid.js'])
              })
              .state('app.table.editable', {
                  url: '/editable',
                  templateUrl: '/greenhome/static/angulr/tpl/table_editable.html',
                  controller: 'XeditableCtrl',
                  resolve: load(['xeditable','/greenhome/static/angulr/js/controllers/xeditable.js'])
              })
              .state('app.table.smart', {
                  url: '/smart',
                  templateUrl: '/greenhome/static/angulr/tpl/table_smart.html',
                  resolve: load(['smart-table','/greenhome/static/angulr/js/controllers/table.js'])
              })
              // form
              .state('app.form', {
                  url: '/form',
                  template: '<div ui-view class="fade-in"></div>',
                  resolve: load('/greenhome/static/angulr/js/controllers/form.js')
              })
              .state('app.form.components', {
                  url: '/components',
                  templateUrl: '/greenhome/static/angulr/tpl/form_components.html',
                  resolve: load(['ngBootstrap','daterangepicker','js/controllers/form.components.js'])
              })
              .state('app.form.elements', {
                  url: '/elements',
                  templateUrl: '/greenhome/static/angulr/tpl/form_elements.html'
              })
              .state('app.form.validation', {
                  url: '/validation',
                  templateUrl: '/greenhome/static/angulr/tpl/form_validation.html'
              })
              .state('app.form.wizard', {
                  url: '/wizard',
                  templateUrl: '/greenhome/static/angulr/tpl/form_wizard.html'
              })
              .state('app.form.fileupload', {
                  url: '/fileupload',
                  templateUrl: 'tpl/form_fileupload.html',
                  resolve: load(['angularFileUpload','/greenhome/static/angulr/js/controllers/file-upload.js'])
              })
              .state('app.form.imagecrop', {
                  url: '/imagecrop',
                  templateUrl: '/greenhome/static/angulr/tpl/form_imagecrop.html',
                  resolve: load(['ngImgCrop','/greenhome/static/angulr/js/controllers/imgcrop.js'])
              })
              .state('app.form.select', {
                  url: '/select',
                  templateUrl: '/greenhome/static/angulr/tpl/form_select.html',
                  controller: 'SelectCtrl',
                  resolve: load(['ui.select','/greenhome/static/angulr/js/controllers/select.js'])
              })
              .state('app.form.slider', {
                  url: '/slider',
                  templateUrl: '/greenhome/static/angulr/tpl/form_slider.html',
                  controller: 'SliderCtrl',
                  resolve: load(['vr.directives.slider','/greenhome/static/angulr/js/controllers/slider.js'])
              })
              .state('app.form.editor', {
                  url: '/editor',
                  templateUrl: '/greenhome/static/angulr/tpl/form_editor.html',
                  controller: 'EditorCtrl',
                  resolve: load(['textAngular','/greenhome/static/angulr/js/controllers/editor.js'])
              })
              .state('app.form.xeditable', {
                  url: '/xeditable',
                  templateUrl: '/greenhome/static/angulr/tpl/form_xeditable.html',
                  controller: 'XeditableCtrl',
                  resolve: load(['xeditable','/greenhome/static/angulr/js/controllers/xeditable.js'])
              })
              // pages
              .state('app.page', {
                  url: '/page',
                  template: '<div ui-view class="fade-in-down"></div>'
              })
              .state('app.page.profile', {
                  url: '/profile',
                  templateUrl: '/greenhome/static/angulr/tpl/page_profile.html'
              })
              .state('app.page.post', {
                  url: '/post',
                  templateUrl: '/greenhome/static/angulr/tpl/page_post.html'
              })
              .state('app.page.search', {
                  url: '/search',
                  templateUrl: '/greenhome/static/angulr/tpl/page_search.html'
              })
              .state('app.page.invoice', {
                  url: '/invoice',
                  templateUrl: '/greenhome/static/angulr/tpl/page_invoice.html'
              })
              .state('app.page.price', {
                  url: '/price',
                  templateUrl: '/greenhome/static/angulr/tpl/page_price.html'
              })
              .state('app.docs', {
                  url: '/docs',
                  templateUrl: '/greenhome/static/angulr/tpl/docs.html'
              })
              // others
              .state('lockme', {
                  url: '/lockme',
                  templateUrl: '/greenhome/static/angulr/tpl/page_lockme.html'
              })
              .state('access', {
                  url: '/access',
                  template: '<div ui-view class="fade-in-right-big smooth"></div>'
              })
              .state('access.signin', {
                  url: '/signin',
                  templateUrl: '/greenhome/static/angulr/tpl/page_signin.html',
                  resolve: load( ['js/controllers/signin.js'] )
              })
              .state('access.signup', {
                  url: '/signup',
                  templateUrl: '/greenhome/static/angulr/tpl/page_signup.html',
                  resolve: load( ['js/controllers/signup.js'] )
              })
              .state('access.forgotpwd', {
                  url: '/forgotpwd',
                  templateUrl: '/greenhome/static/angulr/tpl/page_forgotpwd.html'
              })
              .state('access.404', {
                  url: '/404',
                  templateUrl: '/greenhome/static/angulr/tpl/page_404.html'
              })

              // fullCalendar
              .state('app.calendar', {
                  url: '/calendar',
                  templateUrl: '/greenhome/static/angulr/tpl/app_calendar.html',
                  // use resolve to load other dependences
                  resolve: load(['moment','fullcalendar','ui.calendar','/greenhome/static/angulr/js/app/calendar/calendar.js'])
              })

              // mail
              .state('app.mail', {
                  abstract: true,
                  url: '/mail',
                  templateUrl: 'tpl/mail.html',
                  // use resolve to load other dependences
                  resolve: load( ['/greenhome/static/angulr/js/app/mail/mail.js','/greenhome/static/angulr/js/app/mail/mail-service.js','moment'] )
              })
              .state('app.mail.list', {
                  url: '/inbox/{fold}',
                  templateUrl: '/greenhome/static/angulr/tpl/mail.list.html'
              })
              .state('app.mail.detail', {
                  url: '/{mailId:[0-9]{1,4}}',
                  templateUrl: '/greenhome/static/angulr/tpl/mail.detail.html'
              })
              .state('app.mail.compose', {
                  url: '/compose',
                  templateUrl: '/greenhome/static/angulr/tpl/mail.new.html'
              })

              .state('layout', {
                  abstract: true,
                  url: '/layout',
                  templateUrl: '/greenhome/static/angulr/tpl/layout.html'
              })
              .state('layout.fullwidth', {
                  url: '/fullwidth',
                  views: {
                      '': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_fullwidth.html'
                      },
                      'footer': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_footer_fullwidth.html'
                      }
                  },
                  resolve: load( ['/greenhome/static/angulr/js/controllers/vectormap.js'] )
              })
              .state('layout.mobile', {
                  url: '/mobile',
                  views: {
                      '': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_mobile.html'
                      },
                      'footer': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_footer_mobile.html'
                      }
                  }
              })
              .state('layout.app', {
                  url: '/app',
                  views: {
                      '': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_app.html'
                      },
                      'footer': {
                          templateUrl: '/greenhome/static/angulr/tpl/layout_footer_fullwidth.html'
                      }
                  },
                  resolve: load( ['/greenhome/static/angulr/js/controllers/tab.js'] )
              })
              .state('apps', {
                  abstract: true,
                  url: '/apps',
                  templateUrl: '/greenhome/static/angulr/tpl/layout.html'
              })
              .state('apps.note', {
                  url: '/note',
                  templateUrl: '/greenhome/static/angulr/tpl/apps_note.html',
                  resolve: load( ['/greenhome/static/angulr/js/app/note/note.js','moment'] )
              })
              .state('apps.contact', {
                  url: '/contact',
                  templateUrl: '/greenhome/static/angulr/tpl/apps_contact.html',
                  resolve: load( ['/greenhome/static/angulr/js/app/contact/contact.js'] )
              })
              .state('app.weather', {
                  url: '/weather',
                  templateUrl: '/greenhome/static/angulr/tpl/apps_weather.html',
                  resolve: load(['/greenhome/static/angulr/js/app/weather/skycons.js','angular-skycons','/greenhome/static/angulr/js/app/weather/ctrl.js','moment'])
              })
              .state('app.todo', {
                  url: '/todo',
                  templateUrl: '/greenhome/static/angulr/tpl/apps_todo.html',
                  resolve: load(['/greenhome/static/angulr/js/app/todo/todo.js', 'moment'])
              })
              .state('app.todo.list', {
                  url: '/{fold}'
              })
              .state('app.note', {
                  url: '/note',
                  templateUrl: '/greenhome/static/angulr/tpl/apps_note_material.html',
                  resolve: load(['/greenhome/static/angulr/js/app/note/note.js', 'moment'])
              })
              .state('music', {
                  url: '/music',
                  templateUrl: '/greenhome/static/angulr/tpl/music.html',
                  controller: 'MusicCtrl',
                  resolve: load([
                            'com.2fdevs.videogular', 
                            'com.2fdevs.videogular.plugins.controls', 
                            'com.2fdevs.videogular.plugins.overlayplay',
                            'com.2fdevs.videogular.plugins.poster',
                            'com.2fdevs.videogular.plugins.buffering',
                            '/greenhome/static/angulr/js/app/music/ctrl.js',
                            '/greenhome/static/angulr/js/app/music/theme.css'
                          ])
              })
                  .state('music.home', {
                      url: '/home',
                      templateUrl: '/greenhome/static/angulr/tpl/music.home.html'
                  })
                  .state('music.genres', {
                      url: '/genres',
                      templateUrl: '/greenhome/static/angulr/tpl/music.genres.html'
                  })
                  .state('music.detail', {
                      url: '/detail',
                      templateUrl: '/greenhome/static/angulr/tpl/music.detail.html'
                  })
                  .state('music.mtv', {
                      url: '/mtv',
                      templateUrl: '/greenhome/static/angulr/tpl/music.mtv.html'
                  })
                  .state('music.mtvdetail', {
                      url: '/mtvdetail',
                      templateUrl: '/greenhome/static/angulr/tpl/music.mtv.detail.html'
                  })
                  .state('music.playlist', {
                      url: '/playlist/{fold}',
                      templateUrl: '/greenhome/static/angulr/tpl/music.playlist.html'
                  })
              .state('app.material', {
                  url: '/material',
                  template: '<div ui-view class="wrapper-md"></div>',
                  resolve: load(['/greenhome/static/angulr/js/controllers/material.js'])
                })
                .state('app.material.button', {
                  url: '/button',
                  templateUrl: '/greenhome/static/angulr/tpl/material/button.html'
                })
                .state('app.material.color', {
                  url: '/color',
                  templateUrl: '/greenhome/static/angulr/tpl/material/color.html'
                })
                .state('app.material.icon', {
                  url: '/icon',
                  templateUrl: '/greenhome/static/angulr/tpl/material/icon.html'
                })
                .state('app.material.card', {
                  url: '/card',
                  templateUrl: '/greenhome/static/angulr/tpl/material/card.html'
                })
                .state('app.material.form', {
                  url: '/form',
                  templateUrl: '/greenhome/static/angulr/tpl/material/form.html'
                })
                .state('app.material.list', {
                  url: '/list',
                  templateUrl: '/greenhome/static/angulr/tpl/material/list.html'
                })
                .state('app.material.ngmaterial', {
                  url: '/ngmaterial',
                  templateUrl: '/greenhome/static/angulr/tpl/material/ngmaterial.html'
                });

          function load(srcs, callback) {
            return {
                deps: ['$ocLazyLoad', '$q',
                  function( $ocLazyLoad, $q ){
                    var deferred = $q.defer();
                    var promise  = false;
                    srcs = angular.isArray(srcs) ? srcs : srcs.split(/\s+/);
                    if(!promise){
                      promise = deferred.promise;
                    }
                    angular.forEach(srcs, function(src) {
                      promise = promise.then( function(){
                        if(JQ_CONFIG[src]){
                          return $ocLazyLoad.load(JQ_CONFIG[src]);
                        }
                        angular.forEach(MODULE_CONFIG, function(module) {
                          if( module.name == src){
                            name = module.name;
                          }else{
                            name = src;
                          }
                        });
                        return $ocLazyLoad.load(name);
                      } );
                    });
                    deferred.resolve();
                    return callback ? promise.then(function(){ return callback(); }) : promise;
                }]
            }
          }


      }
    ]
  );
