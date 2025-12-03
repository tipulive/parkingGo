import { redirect } from "@sveltejs/kit";

export function load(){
 throw redirect(301,"you can not access this page")
 
}