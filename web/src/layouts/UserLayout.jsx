import React from 'react';
import { Outlet } from 'react-router-dom';
import Footer from '../components/Footer';
import Header from '../components/Header';

const UserLayout = () => (
  <>
    <Header workspace='user' />
    <div className='main-content router-layout-content'>
      <Outlet />
    </div>
    <Footer />
  </>
);

export default UserLayout;
