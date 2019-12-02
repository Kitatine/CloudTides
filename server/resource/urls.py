from django.conf.urls import url
from django.urls import path, include
from .views import *

urlpatterns = [
    path('validate/', ValidateResource.as_view(), name='validate'),
    path('add/', AddResource.as_view(), name='add')
]