import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HamastarsComponent } from './hamastars/hamastars.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { HamastarDetailComponent } from './hamastar-detail/hamastar-detail.component';


const routes: Routes = [
  { path: 'hamastars', component: HamastarsComponent },
  { path: 'detail/:id', component: HamastarDetailComponent },
  { path:'dashboard', component:DashboardComponent},
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }