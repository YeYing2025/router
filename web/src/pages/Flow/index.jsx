import React from 'react';
import BusinessFlowTable from '../../components/BusinessFlowTable';
import { AppSection } from '../../router-ui';

const FlowPage = ({ kind }) => {
  return (
    <div className='dashboard-container'>
      <AppSection>
          <BusinessFlowTable kind={kind} />
      </AppSection>
    </div>
  );
};

export default FlowPage;
