import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { waitForAsync, ComponentFixture, TestBed } from '@angular/core/testing';

import { DataFeedCreateComponent } from './data-feed-create.component';
import { RouterTestingModule } from '@angular/router/testing';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MockComponent } from 'ng2-mock-component';
import { StoreModule } from '@ngrx/store';
import { ngrxReducers, runtimeChecks } from 'app/ngrx.reducers';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';
import { Regex } from 'app/helpers/auth/regex';
import { Destination } from 'app/entities/destinations/destination.model';

describe('DataFeedCreateComponent', () => {
  let component: DataFeedCreateComponent;
  let fixture: ComponentFixture<DataFeedCreateComponent>;
  // let element;
  let createForm: FormGroup;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [
        MockComponent({ selector: 'chef-th' }),
        MockComponent({ selector: 'chef-td' }),
        MockComponent({ selector: 'chef-error' }),
        MockComponent({ selector: 'chef-form-field' }),
        MockComponent({ selector: 'chef-heading' }),
        MockComponent({ selector: 'chef-icon' }),
        MockComponent({ selector: 'chef-loading-spinner' }),
        MockComponent({ selector: 'mat-select' }),
        MockComponent({ selector: 'mat-option' }),
        MockComponent({ selector: 'chef-page-header' }),
        MockComponent({ selector: 'chef-subheading' }),
        MockComponent({ selector: 'chef-toolbar' }),
        MockComponent({ selector: 'chef-table' }),
        MockComponent({ selector: 'chef-thead' }),
        MockComponent({ selector: 'chef-tbody' }),
        MockComponent({ selector: 'chef-tr' }),
        MockComponent({ selector: 'chef-th' }),
        MockComponent({ selector: 'chef-td' }),
        MockComponent({ selector: 'a', inputs: ['routerLink'] }),
        MockComponent({ selector: 'input', inputs: ['resetOrigin'] }),
        DataFeedCreateComponent
      ],
      providers: [
        FeatureFlagsService
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        RouterTestingModule,
        StoreModule.forRoot(ngrxReducers, { runtimeChecks })
      ],
      schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DataFeedCreateComponent);
    component = fixture.componentInstance;
    component = fixture.componentInstance;
    component.createForm = new FormBuilder().group({
      name: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
      // Note that URL here may be FQDN -or- IP!
      url: ['', [Validators.required, Validators.pattern(Regex.patterns.VALID_FQDN)]],
      tokenType: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
      token: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
      username: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]],
      password: ['', [Validators.required, Validators.pattern(Regex.patterns.NON_BLANK)]]
    });
    createForm = component.createForm;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('Data Feed Create', () => {
    const tokenType = 'Bearer';
    const token = 'test123';
    const destination = <Destination> {
      id: '1',
      name: 'new data feed',
      secret: 'testSecret',
      url: 'http://foo.com'
    };

    it('should be invalid when no fields are filled out', () => {
      expect(createForm.valid).toBeFalsy();
    });

    it('should be valid when all fields are filled out', () => {
      component.integrationSelected = true;
      component.selectChangeHandlers('Access Token');
      component.createForm.controls['name'].setValue(destination.name);
      component.createForm.controls['url'].setValue(destination.url);
      component.createForm.controls['tokenType'].setValue(tokenType);
      component.createForm.controls['token'].setValue(token);
      expect(component.validateForm()).toBeTruthy();
    });

    it('slider resets name, url, username and password to empty string for servicenow', () => {
      component.createForm.controls['name'].setValue('any');
      component.createForm.controls['url'].setValue('any');
      component.createForm.controls['tokenType'].setValue(tokenType);
      component.createForm.controls['token'].setValue('any');
      component.slidePanel();
      component.name = jasmine.createSpyObj('name', ['nativeElement']);
      component.name.nativeElement = { focus: () => { }};
      component.selectIntegration('ServiceNow');
      expect(component.createForm.controls['name'].value).toBe(null);
      expect(component.createForm.controls['url'].value).toBe(null);
      expect(component.createForm.controls['tokenType'].value).toBe(tokenType);
      expect(component.createForm.controls['token'].value).toBe(null);
    });

    it('slider resets name, url, username and password to empty string for splunk', () => {
      component.createForm.controls['name'].setValue('any');
      component.createForm.controls['url'].setValue('any');
      component.createForm.controls['tokenType'].setValue(tokenType);
      component.createForm.controls['token'].setValue('any');
      component.slidePanel();
      component.name = jasmine.createSpyObj('name', ['nativeElement']);
      component.name.nativeElement = { focus: () => { }};
      component.selectIntegration('Splunk');
      expect(component.createForm.controls['name'].value).toBe(null);
      expect(component.createForm.controls['url'].value).toBe(null);
      expect(component.createForm.controls['tokenType'].value).toBe('Splunk');
      expect(component.createForm.controls['token'].value).toBe(null);
    });
  });

  describe('create data feed form validation', () => {

    it('- url field validity', () => {
      component.integrationSelected = true;
      component.name = jasmine.createSpyObj('name', ['nativeElement']);
      component.name.nativeElement = { focus: () => { }};
      component.selectIntegration('ServiceNow');
      expect(component.createForm.controls['url'].value).toBe(null);

      let errors = {};
      const url = component.createForm.controls['url'];
      expect(url.valid).toBeFalsy();

      // url field is required
      errors = url.errors || {};
      expect(errors['required']).toBeTruthy();

      url.setValue('');
      errors = url.errors || {};
      expect(errors['required']).toBeTruthy();

      // Set url to invalid inputs
      url.setValue('  ');
      errors = url.errors || {};
      expect(errors['pattern']).toBeTruthy();
      expect(errors['required']).toBeFalsy();

      // Set url to valid inputs
      url.setValue('any');
      errors = url.errors || {};
      expect(errors['pattern']).toBeTruthy();
    });
  });
});