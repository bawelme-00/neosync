import { withNeosyncContext } from '@/api-only/neosync-context';
import {
  DeleteJobRunRequest,
  GetJobRunRequest,
} from '@/neosync-api-client/mgmt/v1alpha1/job_pb';
import { RequestContext } from '@/shared';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(
  req: NextRequest,
  { params }: RequestContext
): Promise<NextResponse> {
  return withNeosyncContext(async (ctx) => {
    return ctx.jobsClient.getJobRun(
      new GetJobRunRequest({
        id: params.id,
      })
    );
  })(req);
}

export async function DELETE(
  req: NextRequest,
  { params }: RequestContext
): Promise<NextResponse> {
  return withNeosyncContext(async (ctx) => {
    return ctx.jobsClient.deleteJobRun(
      new DeleteJobRunRequest({
        id: params.id,
      })
    );
  })(req);
}
