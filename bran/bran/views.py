import os
import requests
from django.http import JsonResponse

HODOR_URL = os.environ.get('HODOR_URL', 'http://hodor:8080')

def bran_view(request):
    return JsonResponse({
        'service': 'bran',
        'message': 'Bran is watching...',
        'status': 'ok'
    })

def call_hodor(request):
    try:
        response = requests.get(f'{HODOR_URL}/hodr/', timeout=5)
        return JsonResponse({
            'service': 'bran',
            'message': 'Bran called Hodor!',
            'hodor_response': response.json()
        })
    except Exception as e:
        return JsonResponse({'error': str(e)}, status=500)

def health(request):
    return JsonResponse({'status': 'healthy'})