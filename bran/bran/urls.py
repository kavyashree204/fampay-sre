from django.urls import path
from . import views

urlpatterns = [
    path('bran/', views.bran_view),
    path('bran/call-hodor/', views.call_hodor),
    path('health/', views.health),
]