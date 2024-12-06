# Worst_Code-Ever

README - Logique du Code

Introduction

Ce programme est une simulation complexe en Go, combinant des fonctionnalités telles que l'affichage progressif, la manipulation de texte, et des interactions utilisateur. Le code est structuré pour demander une phrase à l'utilisateur, simuler un traitement lent de celle-ci, et analyser les caractères fournis.

Fonctionnalités Principales

Initialisation et Construction d'une Phrase

Une phrase prédéfinie (à l'aide d'un tableau de chaînes) est construite à l'aide d'un strings.Builder. Cela permet de montrer un exemple à l'utilisateur avant qu'il ne saisisse sa propre phrase.

Simulateur de Temps d'Attente

Une goroutine parallèle affiche un message d'attente ("Veuillez patienter, la machine se prépare...") pendant 12 secondes avant de permettre l’avancement du programme.

Lecture de la Phrase de l'Utilisateur

Utilisation de bufio.Scanner pour capturer la saisie utilisateur.

Transformation des caractères saisis en rune pour une meilleure manipulation.

Traitement des Caractères avec Délai Aléatoire

Chaque caractère de la phrase saisie est traité avec un délai aléatoire (entre 1 et 10 secondes).

Si un nombre aléatoire généré est égal à 0, une erreur simulée redémarre le traitement depuis le début.

Structure de Liste Chaînée pour Stocker les Caractères

Chaque caractère est ajouté à une liste chaînée à l’aide d'une structure LetterNode.

Cette structure permet de parcourir et de manipuler les caractères de manière itérative.

Tri des Caractères

Les caractères sont triés à l'aide de l'algorithme de tri à bulles.

Cela assure un affichage alphabétique du contenu de la phrase.

Vérification des Résultats

Le programme compare le nombre total de caractères présents avant et après le tri pour vérifier la cohérence des données.

Une incohérence génère un message d'erreur et arrête le programme.

Affichage du Nombre de Lettres

Le nombre total de lettres est calculé et affiché après un délai pour simuler un traitement "intensif".

Points Notables

Langues et Symboles Multinationaux

Les noms de variables et de structures utilisent différents alphabets (кириллица, japonais, estonien, etc.). Cela n’impacte pas la fonctionnalité du code mais reflète une diversité linguistique.

Goroutines et Communication

Une goroutine est utilisée pour afficher des messages d'attente en parallèle, ce qui introduit un aspect asynchrone dans l'exécution.

Un canal (chan bool) synchronise l'attente.

Simulation de Pannes

Une condition aléatoire simule une "panne de la machine" pour renforcer l’illusion d’un processus complexe.

Tri et Structures Liées

L'utilisation d'une liste chaînée pour les caractères, bien que redondante dans ce cas, montre une approche plus complexe pour gérer les données.

Délais et Aléatoire

Chaque caractère a un délai spécifique en fonction de sa valeur, introduisant une logique réaliste (bien que fictive).