import { TestBed, inject } from '@angular/core/testing';

import { JobLogService, JobLogDefaultService } from './job-log.service';
import { SharedModule } from '../shared/shared.module';
import { SERVICE_CONFIG, IServiceConfig } from '../service.config';

describe('JobLogService', () => {
  const mockConfig: IServiceConfig = {
    replicationJobEndpoint: "/uai-harbor/api/jobs/replication/testing",
    scanJobEndpoint: "/uai-harbor/api/jobs/scan/testing"
  };

  let config: IServiceConfig;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        SharedModule
      ],
      providers: [
        JobLogDefaultService,
        {
          provide: JobLogService,
          useClass: JobLogDefaultService
        }, {
          provide: SERVICE_CONFIG,
          useValue: mockConfig
        }]
    });

    config = TestBed.get(SERVICE_CONFIG);
  });

  it('should be initialized', inject([JobLogDefaultService], (service: JobLogService) => {
    expect(service).toBeTruthy();
    expect(config.replicationJobEndpoint).toEqual("/uai-harbor/api/jobs/replication/testing");
    expect(config.scanJobEndpoint).toEqual("/uai-harbor/api/jobs/scan/testing");
  }));
});
